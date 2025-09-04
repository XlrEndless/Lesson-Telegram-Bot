package port

import (
	"TgBot/cmd/app/output/persist/entity"
	"context"
)

type ITeacherStorage interface {
	GetTeacherById(ctx context.Context, id int) (entity.Teacher, error)
	GetLessonsForStudent(ctx context.Context, student entity.Student) ([]*entity.Lesson, error)
}
