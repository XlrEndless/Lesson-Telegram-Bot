package dto

type OutputMessage[T any] struct {
	ChatId      int    `json:"chat_id"`
	Text        string `json:"text"`
	ReplyMarkup T      `json:"reply_markup"`
}
