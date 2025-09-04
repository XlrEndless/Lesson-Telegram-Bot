package dto

type InputMessage struct {
	MessageId   int                  `json:"message_id"`
	Text        string               `json:"text"`
	From        User                 `json:"from"`
	Chat        Chat                 `json:"chat"`
	Date        int                  `json:"date"`
	ReplyMarkup InlineKeyboardMarkup `json:"reply_markup"`
}
