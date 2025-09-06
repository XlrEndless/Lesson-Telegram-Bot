package dto

import "TgBot/cmd/core/model"

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

func MapUserToModel(input User) model.User {
	return model.User{
		Id:        input.Id,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Username:  input.Username,
	}
}

func MapUserToDto(input model.User) User {
	return User{
		Id:        input.Id,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Username:  input.Username,
	}
}
