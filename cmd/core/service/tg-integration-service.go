package service

import "TgBot/cmd/core/model"

type ITelegramIntegrationService interface {
	GetUpdatesWithOffset(offset int64) ([]model.Update, error)
	SendData(message model.OutputMessage[any]) error
}
