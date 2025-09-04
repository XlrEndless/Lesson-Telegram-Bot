package mapper

import (
	"TgBot/cmd/app/output/persist/entity"
	"TgBot/cmd/core/model"
)

type OffsetMapper struct {
}

func NewOffsetMapper() *OffsetMapper {
	return &OffsetMapper{}
}

func (mapper *OffsetMapper) MapToModel(input entity.Offset) model.Offset {
	output := model.Offset{}
	output.ID = input.ID
	output.Offset = input.Offset
	return output
}

func (mapper *OffsetMapper) MapToEntity(input model.Offset) entity.Offset {
	output := entity.Offset{}
	output.ID = input.ID
	output.Offset = input.Offset
	return output
}
