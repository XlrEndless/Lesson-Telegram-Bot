package entity

import (
	"TgBot/cmd/core/model"
	"gorm.io/gorm"
)

type Placeholder struct {
	gorm.Model
	ID     uint
	Text   string
	UseFor string
}

func MapPlaceholderToModel(input Placeholder) model.Placeholder {
	output := model.Placeholder{}
	output.ID = input.ID
	output.Text = input.Text
	output.UseFor = input.UseFor
	return output
}

func MapPlaceholderToEntity(input model.Placeholder) Placeholder {
	output := Placeholder{}
	output.ID = input.ID
	output.Text = input.Text
	output.UseFor = input.UseFor
	return output
}
