package dto

import "TgBot/cmd/core/model"

type Chat struct {
	ID int `json:"id"`
}

func MapChatToModel(input Chat) model.Chat {
	return model.Chat{
		ID: input.ID,
	}
}

func MapChatToDto(input model.Chat) Chat {
	return Chat{
		ID: input.ID,
	}
}
