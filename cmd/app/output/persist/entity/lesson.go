package entity

import (
	"TgBot/cmd/core/model"
	"gorm.io/gorm"
	"time"
)

type Lesson struct {
	gorm.Model
	ID        uint
	Name      string
	TeacherId int
	Start     time.Time
	End       time.Time
	Students  []Student `gorm:"many2many:student_to_lesson;"`
}

func MapLessonToModel(input Lesson) model.Lesson {
	output := model.Lesson{}
	output.ID = input.ID
	output.Start = input.Start
	output.End = input.End
	output.Name = input.Name
	output.TeacherId = input.TeacherId
	students := make([]model.Student, 0)
	for _, ent := range input.Students {
		student := model.Student{}
		MapStudentToModel(ent, &student)
		students = append(students, student)
	}
	output.Students = students
	return output
}

func MapLessonToEntity(input model.Lesson) Lesson {
	output := Lesson{}
	output.ID = input.ID
	output.Start = input.Start
	output.End = input.End
	output.Name = input.Name
	output.TeacherId = input.TeacherId
	students := make([]Student, 0)
	for _, student := range input.Students {
		ent := Student{}
		MapStudentToEntity(student, &ent)
		students = append(students, ent)
	}
	output.Students = students
	return output
}
