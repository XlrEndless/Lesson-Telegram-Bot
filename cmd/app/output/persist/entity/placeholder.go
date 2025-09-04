package entity

import "gorm.io/gorm"

type Placeholder struct {
	gorm.Model
	ID     uint
	Text   string
	UseFor string
}
