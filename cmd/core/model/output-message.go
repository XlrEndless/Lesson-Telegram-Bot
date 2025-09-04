package model

type OutputMessage[T any] struct {
	ChatId       int
	MessageId    int
	Text         string
	ReplyMarkup  T
	MessageType  string
	MarkupType   string
	CallbackData CallbackData
}

func (outputMessage *OutputMessage[any]) IsEmpty() bool {
	return outputMessage.ChatId == 0
}
