package model

type OutputMessage[T any] struct {
	ChatId       int
	MessageId    int
	Text         string
	ReplyMarkup  T
	MessageType  string
	MarkupType   string
	NeedAnswer   bool
	ShouldDelete bool
	CallbackData CallbackData
}

func (outputMessage *OutputMessage[any]) IsEmpty() bool {
	return outputMessage.ChatId == 0
}
