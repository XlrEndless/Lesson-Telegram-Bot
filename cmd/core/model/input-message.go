package model

type InputMessage struct {
	MessageId   int
	Text        string
	From        User
	Chat        Chat
	Date        int
	ReplyMarkup InlineKeyboardMarkup
	Data        string
}
