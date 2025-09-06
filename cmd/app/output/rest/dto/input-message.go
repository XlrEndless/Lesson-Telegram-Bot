package dto

import "TgBot/cmd/core/model"

type InputMessage struct {
	MessageId   int                  `json:"message_id"`
	Text        string               `json:"text"`
	From        User                 `json:"from"`
	Chat        Chat                 `json:"chat"`
	Date        int                  `json:"date"`
	ReplyMarkup InlineKeyboardMarkup `json:"reply_markup"`
}

func MapInputMessage(input InputMessage) model.InputMessage {
	return model.InputMessage{
		MessageId:   input.MessageId,
		Text:        input.Text,
		From:        MapUserToModel(input.From),
		Chat:        MapChatToModel(input.Chat),
		Date:        input.Date,
		ReplyMarkup: MapInlineKeyboardMarkupToModel(input.ReplyMarkup),
	}
}
