package entity

import (
	"TgBot/cmd/core/model"
	"gorm.io/gorm"
)

type Offset struct {
	gorm.Model
	ID     uint
	Offset int64
}

func MapOffsetToModel(input Offset) model.Offset {
	output := model.Offset{}
	output.ID = input.ID
	output.Offset = input.Offset
	return output
}

func MapOffsetToEntity(input model.Offset) Offset {
	output := Offset{}
	output.ID = input.ID
	output.Offset = input.Offset
	return output
}
