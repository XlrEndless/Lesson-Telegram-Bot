package repository

import (
	"context"
	"gorm.io/gorm"
)

type BaseRepository struct {
	db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) BaseRepository {
	return BaseRepository{db: db}
}

func (repo *BaseRepository) resolveTransaction(ctx context.Context) *gorm.DB {
	value := ctx.Value("db")
	if value == nil {
		return repo.db
	}
	return value.(*gorm.DB)
}
