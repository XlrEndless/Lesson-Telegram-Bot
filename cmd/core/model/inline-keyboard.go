package model

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton
}

type InlineKeyboardButton struct {
	Text         string
	CallbackData string
}
