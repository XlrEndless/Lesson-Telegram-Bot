package dto

import "TgBot/cmd/core/model"

type CallbackQuery struct {
	Id           string       `json:"id"`
	From         User         `json:"from"`
	Data         string       `json:"data"`
	Message      InputMessage `json:"message"`
	ChatInstance string       `json:"chat_instance"`
}

func MapCallbackQuery(input CallbackQuery) model.CallbackQuery {
	return model.CallbackQuery{
		Id:           input.Id,
		From:         MapUserToModel(input.From),
		Data:         input.Data,
		Message:      MapInputMessage(input.Message),
		ChatInstance: input.ChatInstance,
	}
}
