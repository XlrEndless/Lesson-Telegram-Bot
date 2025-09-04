package repository

import (
	"TgBot/cmd/app/output/persist/entity"
	"TgBot/cmd/app/output/persist/mapper"
	"TgBot/cmd/core/model"
	"TgBot/cmd/core/service/port"
	"context"
	"gorm.io/gorm"
)

type OffsetRepository struct {
	BaseRepository
	db           *gorm.DB
	offsetMapper *mapper.OffsetMapper
}

func NewOffsetRepository(
	db *gorm.DB,
	baseRepository BaseRepository,
	offsetMapper *mapper.OffsetMapper) port.IOffsetStorage {
	return &OffsetRepository{baseRepository, db, offsetMapper}
}

func (repo *OffsetRepository) GetFirstOffset(ctx context.Context) (model.Offset, error) {
	db := repo.resolveTransaction(ctx)
	e := entity.Offset{}
	err := db.First(&e).Error
	offset := repo.offsetMapper.MapToModel(e)
	return offset, err
}

func (repo *OffsetRepository) UpdateFirstOffset(ctx context.Context, offset model.Offset) (model.Offset, error) {
	db := repo.resolveTransaction(ctx)
	e := repo.offsetMapper.MapToEntity(offset)
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Save(&e).Error
		return err
	})
	offset = repo.offsetMapper.MapToModel(e)
	return offset, err
}
