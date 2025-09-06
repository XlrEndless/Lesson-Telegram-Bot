package dto

type DeleteMessage struct {
	ChatId    int `json:"chat_id"`
	MessageId int `json:"message_id"`
}
