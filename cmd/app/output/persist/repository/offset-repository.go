package repository

import (
	"TgBot/cmd/app/output/persist/entity"
	"TgBot/cmd/core/model"
	"TgBot/cmd/core/service/port"
	"context"
	"gorm.io/gorm"
)

type OffsetRepository struct {
	BaseRepository
	db *gorm.DB
}

func NewOffsetRepository(
	db *gorm.DB,
	baseRepository BaseRepository,
) port.IOffsetStorage {
	return &OffsetRepository{baseRepository, db}
}

func (repo *OffsetRepository) GetFirstOffset(ctx context.Context) (model.Offset, error) {
	db := repo.resolveTransaction(ctx)
	e := entity.Offset{}
	err := db.First(&e).Error
	offset := entity.MapOffsetToModel(e)
	return offset, err
}

func (repo *OffsetRepository) UpdateFirstOffset(ctx context.Context, offset model.Offset) (model.Offset, error) {
	db := repo.resolveTransaction(ctx)
	e := entity.MapOffsetToEntity(offset)
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Save(&e).Error
		return err
	})
	offset = entity.MapOffsetToModel(e)
	return offset, err
}
