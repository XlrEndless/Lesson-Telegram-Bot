package entity

import "gorm.io/gorm"

type Offset struct {
	gorm.Model
	ID     uint
	Offset int64
}
