package repository

import (
	"TgBot/cmd/app/output/persist/entity"
	"TgBot/cmd/app/output/persist/mapper"
	"TgBot/cmd/core/service/port"
	"context"
	"errors"
	"gorm.io/gorm"
)

type StudentRepository struct {
	BaseRepository
	db            *gorm.DB
	studentMapper *mapper.StudentMapper
}

func NewStudentRepository(
	db *gorm.DB,
	baseRepository BaseRepository,
	studentMapper *mapper.StudentMapper) port.IStudentStorage {
	return &StudentRepository{baseRepository, db, studentMapper}
}

func (repo *StudentRepository) SaveOrUpdateStudent(ctx context.Context, student *entity.Student) error {
	db := repo.resolveTransaction(ctx)
	return db.Transaction(func(tx *gorm.DB) error {
		oldStudent := entity.Student{}
		err := db.Where("tg_id = ?", student.TgId).First(&oldStudent).Error
		if err == nil {
			oldStudent.Name = student.Name
			oldStudent.TgUsername = student.TgUsername
			db.Save(oldStudent)
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = db.Create(student).Error
		}
		return err
	})
}

func (repo *StudentRepository) GetAllStudents(ctx context.Context) ([]entity.Student, error) {
	db := repo.resolveTransaction(ctx)
	var students []entity.Student
	err := db.Find(&students).Error
	return nil, err
}

func (repo *StudentRepository) GetStudentById(ctx context.Context, id int) (entity.Student, error) {
	db := repo.resolveTransaction(ctx)
	var student entity.Student
	err := db.Where("id = ?", id).First(&student).Error
	return student, err
}

func (repo *StudentRepository) GetStudentByTgId(ctx context.Context, id int) (entity.Student, error) {
	db := repo.resolveTransaction(ctx)
	var student entity.Student
	err := db.Where("tg_id = ?", id).First(&student).Error
	return student, err
}
