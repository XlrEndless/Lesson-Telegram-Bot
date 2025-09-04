package service

import (
	"TgBot/cmd/core/model"
	"context"
)

type IResponseBuilder interface {
	BuildResponse(ctx context.Context, strategy string, update model.Update) (model.OutputMessage[any], error)
}
