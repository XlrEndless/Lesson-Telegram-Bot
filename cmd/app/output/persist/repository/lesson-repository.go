package repository

import (
	"TgBot/cmd/app/output/persist/mapper"
	"TgBot/cmd/core/model"
	"TgBot/cmd/core/service/port"
	"context"
	"errors"
	"gorm.io/gorm"
)

type LessonRepository struct {
	BaseRepository
	db           *gorm.DB
	lessonMapper *mapper.LessonMapper
}

func NewLessonRepository(
	db *gorm.DB,
	baseRepository BaseRepository,
	lessonMapper *mapper.LessonMapper) port.ILessonStorage {
	return &LessonRepository{baseRepository, db, lessonMapper}
}

func (repo *LessonRepository) GetById(ctx context.Context, lessonId int) (model.Lesson, error) {
	return model.Lesson{}, errors.New("Not implemented")
}
