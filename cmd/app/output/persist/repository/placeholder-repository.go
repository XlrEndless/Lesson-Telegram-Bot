package repository

import (
	"TgBot/cmd/app/output/persist/entity"
	"TgBot/cmd/core/model"
	"TgBot/cmd/core/service/port"
	"context"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

type PlaceholderRepository struct {
	BaseRepository
	db *gorm.DB
}

func NewPlaceholderRepository(db *gorm.DB, baseRepository BaseRepository) port.IPlaceholderStorage {
	return &PlaceholderRepository{baseRepository, db}
}

func (repo *PlaceholderRepository) GetByUseFor(ctx context.Context, useFor string) (model.Placeholder, error) {
	db := repo.resolveTransaction(ctx)
	e := entity.Placeholder{}
	placeholder := model.Placeholder{}
	err := db.Where("use_for = ?", useFor).First(&e).Error
	if err == nil {
		err = mapstructure.Decode(e, &placeholder)
	}
	return placeholder, err
}

func (repo *PlaceholderRepository) GetAll(ctx context.Context) ([]model.Placeholder, error) {
	db := repo.resolveTransaction(ctx)
	var entities []entity.Placeholder
	err := db.Find(&entities).Error
	var placeholders []model.Placeholder
	if err == nil {
		err = mapstructure.Decode(entities, &placeholders)
	}
	return placeholders, err
}
