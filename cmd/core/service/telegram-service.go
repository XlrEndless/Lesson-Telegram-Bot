package service

import (
	"TgBot/cmd/core/model"
	"context"
)

type ITelegramService interface {
	HandleTelegramUpdate(ctx context.Context, update model.Update) error
	GetUpdates(ctx context.Context) ([]model.Update, error)
	SendResponse(ctx context.Context, update model.OutputMessage[any])
}
