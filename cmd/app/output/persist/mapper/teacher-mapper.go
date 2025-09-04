package mapper

import (
	"TgBot/cmd/app/output/persist/entity"
	"TgBot/cmd/core/model"
)

type TeacherMapper struct {
}

func NewTeacherMapper() *TeacherMapper {
	return &TeacherMapper{}
}

func (mapper *TeacherMapper) MapToModel(input entity.Teacher, output *model.Teacher) *model.Teacher {
	output.ID = input.ID
	output.Name = input.Name
	output.TgId = input.TgId
	output.TgUsername = input.TgUsername
	return output
}

func (mapper *TeacherMapper) MapToEntity(input model.Teacher, output *entity.Teacher) *entity.Teacher {
	output.ID = input.ID
	output.Name = input.Name
	output.TgId = input.TgId
	output.TgUsername = input.TgUsername
	return output
}
