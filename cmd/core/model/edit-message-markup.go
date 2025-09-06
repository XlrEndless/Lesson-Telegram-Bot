package model

type EditMessageReplyMarkup struct {
	BusinessConnectionId string               `json:"business_connection_id"`
	ChatId               int                  `json:"chat_id"`
	MessageId            int                  `json:"message_id"`
	InlineMessageId      string               `json:"inline_message_id"`
	ReplyMarkup          InlineKeyboardMarkup `json:"reply_markup"`
}
