package port

import (
	"TgBot/cmd/app/output/persist/entity"
	"context"
)

type IStudentStorage interface {
	SaveOrUpdateStudent(ctx context.Context, student *entity.Student) error
	GetAllStudents(ctx context.Context) ([]entity.Student, error)
	GetStudentById(ctx context.Context, id int) (entity.Student, error)
	GetStudentByTgId(ctx context.Context, id int) (entity.Student, error)
}
