package dto

import "TgBot/cmd/core/model"

type OutputMessage[T any] struct {
	ChatId      int    `json:"chat_id"`
	Text        string `json:"text"`
	ReplyMarkup T      `json:"reply_markup"`
}

func MapOutputMessageDto(input model.OutputMessage[any]) OutputMessage[any] {
	output := OutputMessage[any]{
		ChatId: input.ChatId,
		Text:   input.Text,
	}
	switch input.ReplyMarkup.(type) {
	case model.InlineKeyboardMarkup:
		output.ReplyMarkup = MapInlineKeyboardMarkupToDto(input.ReplyMarkup.(model.InlineKeyboardMarkup))
	case model.ReplyKeyboardMarkup:
		output.ReplyMarkup = MapReplyKeyboardMarkupToDto(input.ReplyMarkup.(model.ReplyKeyboardMarkup))
	default:
		output.ReplyMarkup = input.ReplyMarkup
	}
	return output
}
