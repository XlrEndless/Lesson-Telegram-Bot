package mapper

import (
	"TgBot/cmd/app/output/persist/entity"
	"TgBot/cmd/core/model"
)

type PlaceholderMapper struct {
}

func NewPlaceholderMapper() *PlaceholderMapper {
	return &PlaceholderMapper{}
}

func (mapper *PlaceholderMapper) MapToModel(input entity.Placeholder) model.Placeholder {
	output := model.Placeholder{}
	output.ID = input.ID
	output.Text = input.Text
	output.UseFor = input.UseFor
	return output
}

func (mapper *PlaceholderMapper) MapToEntity(input model.Placeholder) entity.Placeholder {
	output := entity.Placeholder{}
	output.ID = input.ID
	output.Text = input.Text
	output.UseFor = input.UseFor
	return output
}
