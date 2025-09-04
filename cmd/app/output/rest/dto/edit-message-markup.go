package dto

type EditMessageReplyMarkup struct {
	BusinessConnectionId string `json:"business_connection_id"`
	ChatId               int    `json:"chat_id"`
	MessageId            int    `json:"message_id"`
	InlineMessageId      string `json:"inline_message_id"`
	ReplyMarkup          any    `json:"reply_markup"`
}
