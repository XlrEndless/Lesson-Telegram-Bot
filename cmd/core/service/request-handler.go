package service

import (
	"TgBot/cmd/core/model"
	"context"
)

type IRequestHandler interface {
	HandleRequest(ctx context.Context, update model.Update) (model.OutputMessage[any], error)
}
