package entity

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	ID         uint
	Name       string
	TgId       string
	TgUsername string
}
