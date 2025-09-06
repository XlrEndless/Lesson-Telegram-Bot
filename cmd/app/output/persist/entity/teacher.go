package entity

import (
	"TgBot/cmd/core/model"
	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	ID         uint
	Name       string
	TgId       string
	TgUsername string
}

func MapTeacherToModel(input Teacher) model.Teacher {
	output := model.Teacher{}
	output.ID = input.ID
	output.Name = input.Name
	output.TgId = input.TgId
	output.TgUsername = input.TgUsername
	return output
}

func MapTeacherToEntity(input model.Teacher) Teacher {
	output := Teacher{}
	output.ID = input.ID
	output.Name = input.Name
	output.TgId = input.TgId
	output.TgUsername = input.TgUsername
	return output
}
