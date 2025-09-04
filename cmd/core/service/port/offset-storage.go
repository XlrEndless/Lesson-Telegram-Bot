package port

import (
	"TgBot/cmd/core/model"
	"context"
)

type IOffsetStorage interface {
	GetFirstOffset(ctx context.Context) (model.Offset, error)
	UpdateFirstOffset(ctx context.Context, offset model.Offset) (model.Offset, error)
}
