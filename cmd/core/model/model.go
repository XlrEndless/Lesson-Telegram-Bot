package model

import (
	"time"
)

type Lesson struct {
	ID        uint
	Name      string
	TeacherId int
	Start     time.Time
	End       time.Time
	Students  []Student `gorm:"many2many:student_to_lesson;"`
}

type Offset struct {
	ID     uint
	Offset int64
}

type Student struct {
	ID         uint
	Name       string
	TgId       int
	TgUsername string
	Lessons    []Lesson `gorm:"many2many:student_to_lesson;"`
}

type Teacher struct {
	ID         uint
	Name       string
	TgId       string
	TgUsername string
}

type Placeholder struct {
	ID     uint
	Text   string
	UseFor string
}
