package impl

import (
	"TgBot/cmd/core/model"
	"TgBot/cmd/core/service"
	"TgBot/cmd/core/service/port"
	"context"
	"errors"
	"gorm.io/gorm"
)

type OffsetService struct {
	offsetStorage port.IOffsetStorage
}

func NewOffsetService(offsetStorage port.IOffsetStorage) service.IOffsetService {
	return &OffsetService{offsetStorage}
}

func (offsetService *OffsetService) UpdateOffset(ctx context.Context, offset model.Offset) (model.Offset, error) {
	oldOffsetEntity, err := offsetService.offsetStorage.GetFirstOffset(ctx)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		oldOffsetEntity = model.Offset{Offset: -1}
	}
	if oldOffsetEntity.Offset <= offset.Offset {
		offset.Offset += 1
		oldOffsetEntity.Offset = offset.Offset
		_, err := offsetService.offsetStorage.UpdateFirstOffset(ctx, oldOffsetEntity)
		return offset, err
	}
	offset.Offset = oldOffsetEntity.Offset
	return offset, err
}

func (offsetService *OffsetService) GetOffset(ctx context.Context) (model.Offset, error) {
	oldOffset, err := offsetService.offsetStorage.GetFirstOffset(ctx)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		oldOffset = model.Offset{Offset: -1}
	}
	return oldOffset, err
}
