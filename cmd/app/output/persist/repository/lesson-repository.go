package repository

import (
	"TgBot/cmd/core/model"
	"TgBot/cmd/core/service/port"
	"context"
	"errors"
	"gorm.io/gorm"
)

type LessonRepository struct {
	BaseRepository
	db *gorm.DB
}

func NewLessonRepository(
	db *gorm.DB,
	baseRepository BaseRepository,
) port.ILessonStorage {
	return &LessonRepository{baseRepository, db}
}

func (repo *LessonRepository) GetById(ctx context.Context, lessonId int) (model.Lesson, error) {
	return model.Lesson{}, errors.New("Not implemented")
}
