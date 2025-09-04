package mapper

import (
	"TgBot/cmd/app/output/persist/entity"
	"TgBot/cmd/core/model"
)

type StudentMapper struct {
	lessonMapper *LessonMapper
}

func NewStudentMapper(lessonMapper *LessonMapper) *StudentMapper {
	return &StudentMapper{lessonMapper}
}

func (mapper *StudentMapper) MapToModel(input entity.Student, output *model.Student) *model.Student {
	output.ID = input.ID
	output.Name = input.Name
	output.TgUsername = input.TgUsername
	output.TgId = input.TgId
	lessons := make([]model.Lesson, 0)
	for _, ent := range input.Lessons {
		lesson := model.Lesson{}
		mapper.lessonMapper.MapToModel(ent, &lesson)
		lessons = append(lessons, lesson)
	}
	output.Lessons = lessons
	return output
}

func (mapper *StudentMapper) MapToEntity(input model.Student, output *entity.Student) *entity.Student {
	output.ID = input.ID
	output.Name = input.Name
	output.TgUsername = input.TgUsername
	output.TgId = input.TgId
	lessons := make([]entity.Lesson, 0)
	for _, lesson := range input.Lessons {
		ent := entity.Lesson{}
		mapper.lessonMapper.MapToEntity(lesson, &ent)
		lessons = append(lessons, ent)
	}
	output.Lessons = lessons
	return output
}
