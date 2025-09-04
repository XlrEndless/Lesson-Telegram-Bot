package service

import (
	"TgBot/cmd/core/model"
	"context"
)

type IOffsetService interface {
	UpdateOffset(ctx context.Context, offset model.Offset) (model.Offset, error)
	GetOffset(ctx context.Context) (model.Offset, error)
}
