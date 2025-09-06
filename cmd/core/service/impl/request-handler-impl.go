package impl

import (
	"TgBot/cmd/app/output/persist/entity"
	"TgBot/cmd/core/constant"
	"TgBot/cmd/core/model"
	"TgBot/cmd/core/service"
	"TgBot/cmd/core/service/port"
	"TgBot/cmd/core/tools"
	"context"
	"log/slog"
	"strings"
)

type RequestHandler struct {
	strategies         map[string]func(ctx context.Context, update model.Update) (model.OutputMessage[any], error)
	studentStorage     port.IStudentStorage
	responseBuilder    service.IResponseBuilder
	transactionManager port.ITransactionManager
}

func NewRequestHandler(
	studentStorage port.IStudentStorage,
	responseBuilder service.IResponseBuilder,
	transactionManager port.ITransactionManager) service.IRequestHandler {
	strategies := make(map[string]func(ctx context.Context, update model.Update) (model.OutputMessage[any], error))
	requestHandler := RequestHandler{
		strategies,
		studentStorage,
		responseBuilder,
		transactionManager,
	}
	strategies["/start"] = requestHandler.HandleStartRequest
	strategies["/newStudent"] = requestHandler.HandleAddStudentRequest
	strategies["/dates"] = requestHandler.HandleTimetableRequest
	strategies[constant.RegToLessonLabel] = requestHandler.HandleTimetableRequest
	return &requestHandler
}

func (handler *RequestHandler) HandleRequest(ctx context.Context, update model.Update) (model.OutputMessage[any], error) {
	var err error
	var message model.OutputMessage[any]
	if update.IsCommand() {
		message, err = handler.HandleCommand(ctx, update)
	} else if update.IsQuery() {
		message, err = handler.handleQuery(ctx, update)
	} else if update.IsNotPrefixedCommand() {
		message, err = handler.handleNotPrefixedCommand(ctx, update)
	} else {
		err = tools.StrategyNotFoundErr
	}
	return message, err
}

func (handler *RequestHandler) HandleCommand(ctx context.Context, update model.Update) (model.OutputMessage[any], error) {
	slog.Info("Got command...")
	msgText := update.Message.Text
	update.Message.Data = msgText
	msgText = strings.TrimSpace(msgText)
	strategies := handler.strategies
	strategy, ok := strategies[msgText]
	var err error
	var message model.OutputMessage[any]
	if ok {
		message, err = strategy(ctx, update)
	} else {
		err = tools.StrategyNotFoundErr
	}
	message.MessageType = constant.New
	return message, err
}

func (handler *RequestHandler) handleQuery(ctx context.Context, update model.Update) (model.OutputMessage[any], error) {
	slog.Info("Got callback query...")
	query := update.CallbackQuery
	update.Message = query.Message
	callBackData := query.Data
	update.Message.Data = callBackData
	commands := strings.Split(callBackData, "/")
	strategies := handler.strategies
	strategy, ok := strategies["/"+commands[1]]
	var err error
	var message model.OutputMessage[any]
	if ok {
		message, err = strategy(ctx, update)
		if err == nil {
			message.NeedAnswer = true
			message.MessageId = update.Message.MessageId
			message.CallbackData = model.CallbackData{CallbackId: query.Id}
			if message.MarkupType == constant.Inline {
				message.MessageType = constant.Change
			} else {
				message.ShouldDelete = true
				message.MessageType = constant.New
			}
		}
	} else {
		err = tools.StrategyNotFoundErr
	}
	return message, err
}

func (handler *RequestHandler) handleNotPrefixedCommand(ctx context.Context, update model.Update) (model.OutputMessage[any], error) {
	slog.Info("Got not prefixed command...")
	return handler.HandleCommand(ctx, update)
}

func (handler *RequestHandler) HandleAddStudentRequest(ctx context.Context, update model.Update) (model.OutputMessage[any], error) {
	slog.Info("Handle add student request...")
	err := handler.transactionManager.DoInTransaction(ctx, func(ctx context.Context) error {
		studentRepo := handler.studentStorage
		msg := update.Message
		from := msg.From
		newStudent := entity.Student{TgId: from.Id, Name: from.FirstName, TgUsername: from.Username}
		slog.Info("Save student...")
		return studentRepo.SaveOrUpdateStudent(ctx, &newStudent)
	})
	return model.OutputMessage[any]{}, err
}

func (handler *RequestHandler) HandleStartRequest(ctx context.Context, update model.Update) (model.OutputMessage[any], error) {
	slog.Info("Handle start bot request...")
	return handler.responseBuilder.BuildResponse(ctx, "start", update)
}

func (handler *RequestHandler) HandleTimetableRequest(ctx context.Context, update model.Update) (model.OutputMessage[any], error) {
	slog.Info("Handle get timetable request...")
	return handler.responseBuilder.BuildResponse(ctx, "dates", update)
}
