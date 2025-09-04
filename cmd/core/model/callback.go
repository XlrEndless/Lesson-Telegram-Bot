package model

type CallbackQuery struct {
	Id           string
	From         User
	Data         string
	Message      InputMessage
	ChatInstance string
}

type CallbackData struct {
	CallbackId string
}

type Page struct {
	Offset int
	Limit  int
}
