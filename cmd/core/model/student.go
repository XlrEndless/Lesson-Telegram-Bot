package model

type Student struct {
	ID         uint
	Name       string
	TgId       int
	TgUsername string
	Lessons    []Lesson `gorm:"many2many:student_to_lesson;"`
}
