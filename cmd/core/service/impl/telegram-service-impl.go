package impl

import (
	"TgBot/cmd/core/model"
	"TgBot/cmd/core/service"
	"TgBot/cmd/core/service/port"
	"TgBot/cmd/core/tools"
	"context"
	"errors"
)

type TelegramService struct {
	requestHandler     service.IRequestHandler
	integrationService service.ITelegramIntegrationService
	offsetService      service.IOffsetService
	transactionManager port.ITransactionManager
}

func NewTelegramService(
	integrationService service.ITelegramIntegrationService,
	requestHandler service.IRequestHandler,
	offsetService service.IOffsetService,
	transactionManager port.ITransactionManager) service.ITelegramService {
	return &TelegramService{
		requestHandler,
		integrationService,
		offsetService,
		transactionManager,
	}
}

func (telegramService *TelegramService) GetUpdates(ctx context.Context) ([]model.Update, error) {
	offset, err := telegramService.offsetService.GetOffset(ctx)
	if err != nil {
		return make([]model.Update, 0), err
	}
	updates, err := telegramService.integrationService.GetUpdatesWithOffset(offset.Offset)
	if err != nil {
		return make([]model.Update, 0), err
	}
	return updates, err
}

func (telegramService *TelegramService) SendResponse(ctx context.Context, update model.OutputMessage[any]) {

}

func (telegramService *TelegramService) HandleTelegramUpdate(ctx context.Context, update model.Update) error {
	var outputData model.OutputMessage[any]
	var err error
	requestHandler := telegramService.requestHandler
	outputData, err = requestHandler.HandleRequest(ctx, update)
	if err == nil && !outputData.IsEmpty() {
		err = telegramService.integrationService.SendData(outputData)
	}
	if err == nil || errors.Is(err, tools.StrategyNotFoundErr) {
		offset := model.Offset{Offset: int64(update.ID)}
		_, err = telegramService.offsetService.UpdateOffset(ctx, offset)
	}
	return err
}
