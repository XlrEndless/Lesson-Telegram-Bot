package port

import (
	"TgBot/cmd/core/model"
	"context"
)

type IPlaceholderStorage interface {
	GetByUseFor(ctx context.Context, useFor string) (model.Placeholder, error)
	GetAll(ctx context.Context) ([]model.Placeholder, error)
}
