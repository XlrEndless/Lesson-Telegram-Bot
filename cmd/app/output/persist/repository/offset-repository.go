package repository

import (
	"TgBot/cmd/app/output/persist/entity"
	"TgBot/cmd/core/model"
	"TgBot/cmd/core/service/port"
	"context"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

type OffsetRepository struct {
	BaseRepository
	db *gorm.DB
}

func NewOffsetRepository(db *gorm.DB, baseRepository BaseRepository) port.IOffsetStorage {
	return &OffsetRepository{baseRepository, db}
}

func (repo *OffsetRepository) GetFirstOffset(ctx context.Context) (model.Offset, error) {
	db := repo.resolveTransaction(ctx)
	e := entity.Offset{}
	err := db.First(&e).Error
	offset := model.Offset{}
	if err == nil {
		err = mapstructure.Decode(e, &offset)
	}
	return offset, err
}

func (repo *OffsetRepository) UpdateFirstOffset(ctx context.Context, offset model.Offset) (model.Offset, error) {
	db := repo.resolveTransaction(ctx)
	e := entity.Offset{}
	err := mapstructure.Decode(offset, &e)
	if err != nil {
		return offset, err
	}
	err = db.Transaction(func(tx *gorm.DB) error {
		err := tx.Save(&e).Error
		return err
	})
	if err == nil {
		err = mapstructure.Decode(e, &offset)
	}
	return offset, err
}
