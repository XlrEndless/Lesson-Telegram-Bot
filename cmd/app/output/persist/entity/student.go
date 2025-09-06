package entity

import (
	"TgBot/cmd/core/model"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	ID         uint
	Name       string
	TgId       int
	TgUsername string
	Lessons    []Lesson `gorm:"many2many:student_to_lesson;"`
}

func MapStudentToModel(input Student, output *model.Student) *model.Student {
	output.ID = input.ID
	output.Name = input.Name
	output.TgUsername = input.TgUsername
	output.TgId = input.TgId
	lessons := make([]model.Lesson, 0)
	for _, ent := range input.Lessons {
		lesson := MapLessonToModel(ent)
		lessons = append(lessons, lesson)
	}
	output.Lessons = lessons
	return output
}

func MapStudentToEntity(input model.Student, output *Student) *Student {
	output.ID = input.ID
	output.Name = input.Name
	output.TgUsername = input.TgUsername
	output.TgId = input.TgId
	lessons := make([]Lesson, 0)
	for _, lesson := range input.Lessons {
		ent := MapLessonToEntity(lesson)
		lessons = append(lessons, ent)
	}
	output.Lessons = lessons
	return output
}
