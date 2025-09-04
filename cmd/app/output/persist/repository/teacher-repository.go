package repository

import (
	"TgBot/cmd/app/output/persist/entity"
	"TgBot/cmd/app/output/persist/mapper"
	"TgBot/cmd/core/service/port"
	"context"
	"gorm.io/gorm"
)

type TeacherRepository struct {
	BaseRepository
	db            *gorm.DB
	teacherMapper *mapper.TeacherMapper
}

func NewTeacherRepository(
	db *gorm.DB,
	baseRepository BaseRepository,
	teacherMapper *mapper.TeacherMapper) port.ITeacherStorage {
	return &TeacherRepository{baseRepository, db, teacherMapper}
}

func (repo *TeacherRepository) GetTeacherById(ctx context.Context, id int) (entity.Teacher, error) {
	db := repo.resolveTransaction(ctx)
	var teacher entity.Teacher
	err := db.Where("id = ?", id).First(&teacher).Error
	return teacher, err
}

func (repo *TeacherRepository) GetLessonsForStudent(ctx context.Context, student entity.Student) ([]*entity.Lesson, error) {
	db := repo.resolveTransaction(ctx)
	var lessons []*entity.Lesson
	err := db.Where("student_id = ?", student.ID).Find(&lessons).Error
	return lessons, err
}
