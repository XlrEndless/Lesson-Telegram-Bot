package mapper

import (
	"TgBot/cmd/app/output/persist/entity"
	"TgBot/cmd/core/model"
)

type LessonMapper struct {
	studentMapper *StudentMapper
}

func NewLessonMapper(studentMapper *StudentMapper) *LessonMapper {
	return &LessonMapper{studentMapper}
}

func (mapper *LessonMapper) MapToModel(input entity.Lesson, output *model.Lesson) *model.Lesson {
	output.ID = input.ID
	output.Start = input.Start
	output.End = input.End
	output.Name = input.Name
	output.TeacherId = input.TeacherId
	students := make([]model.Student, 0)
	for _, ent := range input.Students {
		student := model.Student{}
		mapper.studentMapper.MapToModel(ent, &student)
		students = append(students, student)
	}
	output.Students = students
	return output
}

func (mapper *LessonMapper) MapToEntity(input model.Lesson, output *entity.Lesson) *entity.Lesson {
	output.ID = input.ID
	output.Start = input.Start
	output.End = input.End
	output.Name = input.Name
	output.TeacherId = input.TeacherId
	students := make([]entity.Student, 0)
	for _, student := range input.Students {
		ent := entity.Student{}
		mapper.studentMapper.MapToEntity(student, &ent)
		students = append(students, ent)
	}
	output.Students = students
	return output
}
