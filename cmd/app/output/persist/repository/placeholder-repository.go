package repository

import (
	"TgBot/cmd/app/output/persist/entity"
	"TgBot/cmd/core/model"
	"TgBot/cmd/core/service/port"
	"context"
	"gorm.io/gorm"
)

type PlaceholderRepository struct {
	BaseRepository
	db *gorm.DB
}

func NewPlaceholderRepository(
	db *gorm.DB,
	baseRepository BaseRepository) port.IPlaceholderStorage {
	return &PlaceholderRepository{baseRepository, db}
}

func (repo *PlaceholderRepository) GetByUseFor(ctx context.Context, useFor string) (model.Placeholder, error) {
	db := repo.resolveTransaction(ctx)
	e := entity.Placeholder{}
	err := db.Where("use_for = ?", useFor).First(&e).Error
	placeholder := entity.MapPlaceholderToModel(e)
	return placeholder, err
}

func (repo *PlaceholderRepository) GetAll(ctx context.Context) ([]model.Placeholder, error) {
	db := repo.resolveTransaction(ctx)
	var entities []entity.Placeholder
	err := db.Find(&entities).Error
	placeholders := make([]model.Placeholder, 0)
	for _, ent := range entities {
		placeholders = append(placeholders, entity.MapPlaceholderToModel(ent))
	}
	return placeholders, err
}
