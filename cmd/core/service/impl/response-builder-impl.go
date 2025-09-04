package impl

import (
	"TgBot/cmd/core/constant"
	"TgBot/cmd/core/model"
	"TgBot/cmd/core/service"
	"TgBot/cmd/core/service/port"
	"TgBot/cmd/core/tools"
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type ResponseBuilder struct {
	strategies           map[string]func(ctx context.Context, update model.Update) (model.OutputMessage[any], error)
	placeholderStorage   port.IPlaceholderStorage
	defaultDatePageLimit int
	defaultTimePageLimit int
}

func NewResponseBuilder(placeholderStorage port.IPlaceholderStorage) service.IResponseBuilder {
	strategies := map[string]func(ctx context.Context, update model.Update) (model.OutputMessage[any], error){}
	responseBuilder := ResponseBuilder{
		strategies,
		placeholderStorage,
		constant.DefaultDatePageLimit,
		constant.DefaultTimePageLimit,
	}
	strategies["dates"] = responseBuilder.BuildDateTable
	strategies["start"] = responseBuilder.BuildStartResponse
	return &responseBuilder
}

func (builder *ResponseBuilder) BuildResponse(
	ctx context.Context,
	strategy string,
	update model.Update) (model.OutputMessage[any], error) {
	fn := builder.strategies[strategy]
	return fn(ctx, update)
}

func (builder *ResponseBuilder) BuildDateTable(
	ctx context.Context,
	update model.Update) (model.OutputMessage[any], error) {
	message := update.Message
	chat := message.Chat
	var outputMessage model.OutputMessage[any]
	var err error
	page := model.Page{}
	if builder.isPageable(message.Data) {
		page, err = builder.resolvePage(message.Data)
	}
	if err != nil {
		if errors.Is(err, tools.PageNotFoundErr) {
			page = model.Page{Offset: 0}
			err = nil
		} else {
			return outputMessage, err
		}
	}
	if page.Limit == 0 {
		page.Limit = builder.defaultDatePageLimit
	}
	inlineKeyboard := builder.createDateTableButtons(ctx, page.Offset, page.Limit)
	inlineButtonNext := model.InlineKeyboardButton{
		Text:         constant.NextPageLabel,
		CallbackData: fmt.Sprintf("/dates/page{%d}", page.Offset+1),
	}

	inlineKeyboard = append(inlineKeyboard, []model.InlineKeyboardButton{inlineButtonNext})

	if page.Offset > 0 {
		inlineButtonPrev := model.InlineKeyboardButton{
			Text:         constant.PreviousCommandLabel,
			CallbackData: fmt.Sprintf("/dates/page{%d}", page.Offset-1),
		}
		inlineKeyboard = append(inlineKeyboard, []model.InlineKeyboardButton{inlineButtonPrev})
	} else {
		inlineButtonHome := model.InlineKeyboardButton{
			Text:         constant.PreviousCommandLabel,
			CallbackData: "/start",
		}
		inlineKeyboard = append(inlineKeyboard, []model.InlineKeyboardButton{inlineButtonHome})
	}

	replyMarkup := model.InlineKeyboardMarkup{InlineKeyboard: inlineKeyboard}
	outputMessage = model.OutputMessage[any]{
		ChatId:      chat.ID,
		Text:        constant.ChooseDateLabel,
		ReplyMarkup: replyMarkup,
		MarkupType:  constant.Inline,
	}
	return outputMessage, err
}

func (builder *ResponseBuilder) BuildStartResponse(
	ctx context.Context,
	update model.Update) (model.OutputMessage[any], error) {
	message := update.Message
	var err error
	replyKeyboard := make([][]model.ReplyKeyboardButton, 0)
	replyButton1 := model.ReplyKeyboardButton{Text: constant.RegToLessonLabel}
	replyButton := model.ReplyKeyboardButton{Text: constant.ShowMyLessonsLabel}
	replyKeyboard = append(replyKeyboard, []model.ReplyKeyboardButton{replyButton1})
	replyKeyboard = append(replyKeyboard, []model.ReplyKeyboardButton{replyButton})
	replyMarkup := model.ReplyKeyboardMarkup{
		ReplyKeyboardButton: replyKeyboard,
		Selective:           true,
		ResizeKeyboard:      true,
		OneTimeKeyboard:     true,
	}
	chat := message.Chat
	outputMessage := model.OutputMessage[any]{
		ChatId:      chat.ID,
		ReplyMarkup: replyMarkup,
		Text:        constant.ChooseAction,
		MarkupType:  constant.Reply,
	}
	return outputMessage, err
}

func (builder *ResponseBuilder) createDateTableButtons(
	ctx context.Context,
	offset int,
	limit int) [][]model.InlineKeyboardButton {
	fromDate := time.Now().AddDate(0, 0, offset*limit)
	toDate := fromDate.AddDate(0, 0, limit)
	inlineKeyboard := make([][]model.InlineKeyboardButton, 0)
	for fromDate.Before(toDate) {
		date := fromDate.Format(time.DateOnly)
		inlineButton := model.InlineKeyboardButton{
			Text:         date,
			CallbackData: "/showTime",
		}
		fromDate = fromDate.AddDate(0, 0, 1)
		inlineKeyboard = append(inlineKeyboard, []model.InlineKeyboardButton{inlineButton})
	}
	return inlineKeyboard
}

func (builder *ResponseBuilder) isPageable(data string) bool {
	return strings.Contains(data, "/page")
}

func (builder *ResponseBuilder) resolvePage(data string) (model.Page, error) {
	commands := strings.Split(data, "/")
	for _, command := range commands {
		if strings.Contains(command, "page") {
			pageNumber := strings.TrimLeft(strings.TrimRight(command, "}"), "page{")
			page, err := strconv.ParseInt(pageNumber, 10, 64)
			if err != nil {
				return model.Page{}, err
			} else {
				return model.Page{Offset: int(page)}, err
			}
		}
	}
	return model.Page{}, tools.PageNotFoundErr
}

func (builder *ResponseBuilder) buildPlaceholdersMap(
	ctx context.Context,
	placeholders []model.Placeholder,
) (map[string]model.Placeholder, error) {
	placeholders, err := builder.placeholderStorage.GetAll(ctx)
	placeholdersMap := make(map[string]model.Placeholder)
	for _, placeholder := range placeholders {
		placeholdersMap[placeholder.UseFor] = placeholder
	}
	return placeholdersMap, err
}
