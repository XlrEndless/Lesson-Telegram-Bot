package port

import (
	"TgBot/cmd/core/model"
	"context"
)

type ILessonStorage interface {
	GetById(ctx context.Context, lessonId int) (model.Lesson, error)
	//GetLessonByStudentID(id int64) (*entity.Lesson, error)
	//GetLessonByStudentIdForTime(id int64, from time.Time, to time.Time) (*entity.Lesson, error)
	//GetLessonsByTeacherID(id int64) ([]*entity.Lesson, error)
	//GetLessonsByTeacherIdForTime(id int64, from time.Time, to time.Time) ([]*entity.Lesson, error)
	//SaveLesson(lesson *entity.Lesson) (*entity.Lesson, error)
	//UpdateLesson(lesson *entity.Lesson) (*entity.Lesson, error)
}
