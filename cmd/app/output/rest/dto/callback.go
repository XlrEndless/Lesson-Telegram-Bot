package dto

type CallbackQuery struct {
	Id           string       `json:"id"`
	From         User         `json:"from"`
	Data         string       `json:"data"`
	Message      InputMessage `json:"message"`
	ChatInstance string       `json:"chat_instance"`
}
