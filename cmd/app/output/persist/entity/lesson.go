package entity

import (
	"gorm.io/gorm"
	"time"
)

type Lesson struct {
	gorm.Model
	ID       uint
	Name     string
	Teacher  int
	Start    time.Time
	End      time.Time
	Students []*Student `gorm:"many2many:student_to_lesson;"`
}
