package dto

import "TgBot/cmd/core/model"

type EditMessageReplyMarkup struct {
	BusinessConnectionId string               `json:"business_connection_id"`
	ChatId               int                  `json:"chat_id"`
	MessageId            int                  `json:"message_id"`
	InlineMessageId      string               `json:"inline_message_id"`
	ReplyMarkup          InlineKeyboardMarkup `json:"reply_markup"`
}

func MapEditMessageReplyMarkupToModel(input EditMessageReplyMarkup) model.EditMessageReplyMarkup {
	return model.EditMessageReplyMarkup{
		BusinessConnectionId: input.BusinessConnectionId,
		ChatId:               input.ChatId,
		MessageId:            input.MessageId,
		InlineMessageId:      input.InlineMessageId,
		ReplyMarkup:          MapInlineKeyboardMarkupToModel(input.ReplyMarkup),
	}
}

func MapEditMessageReplyMarkupToDto(input model.EditMessageReplyMarkup) EditMessageReplyMarkup {
	return EditMessageReplyMarkup{
		BusinessConnectionId: input.BusinessConnectionId,
		ChatId:               input.ChatId,
		MessageId:            input.MessageId,
		InlineMessageId:      input.InlineMessageId,
		ReplyMarkup:          MapInlineKeyboardMarkupToDto(input.ReplyMarkup),
	}
}
