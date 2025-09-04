package repository

import (
	"TgBot/cmd/app/output/persist/entity"
	"TgBot/cmd/app/output/persist/mapper"
	"TgBot/cmd/core/model"
	"TgBot/cmd/core/service/port"
	"context"
	"gorm.io/gorm"
)

type PlaceholderRepository struct {
	BaseRepository
	db                *gorm.DB
	placeholderMapper *mapper.PlaceholderMapper
}

func NewPlaceholderRepository(
	db *gorm.DB,
	baseRepository BaseRepository,
	placeholderMapper *mapper.PlaceholderMapper) port.IPlaceholderStorage {
	return &PlaceholderRepository{baseRepository, db, placeholderMapper}
}

func (repo *PlaceholderRepository) GetByUseFor(ctx context.Context, useFor string) (model.Placeholder, error) {
	db := repo.resolveTransaction(ctx)
	e := entity.Placeholder{}
	err := db.Where("use_for = ?", useFor).First(&e).Error
	placeholder := repo.placeholderMapper.MapToModel(e)
	return placeholder, err
}

func (repo *PlaceholderRepository) GetAll(ctx context.Context) ([]model.Placeholder, error) {
	db := repo.resolveTransaction(ctx)
	var entities []entity.Placeholder
	err := db.Find(&entities).Error
	placeholders := make([]model.Placeholder, len(entities))
	for _, ent := range entities {
		placeholders = append(placeholders, repo.placeholderMapper.MapToModel(ent))
	}
	return placeholders, err
}
