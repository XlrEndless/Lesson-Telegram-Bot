package model

import "time"

type Lesson struct {
	ID        uint
	Name      string
	TeacherId int
	Start     time.Time
	End       time.Time
	Students  []Student `gorm:"many2many:student_to_lesson;"`
}
