package entity

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	ID         uint
	Name       string
	TgId       int
	TgUsername string
	Lessons    []*Lesson `gorm:"many2many:student_to_lesson;"`
}
