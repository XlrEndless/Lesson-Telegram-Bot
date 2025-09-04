package dto

type ReplyKeyboardMarkup struct {
	ReplyKeyboardButton   [][]ReplyKeyboardButton `json:"keyboard"`
	ResizeKeyboard        bool                    `json:"resize_keyboard"`
	OneTimeKeyboard       bool                    `json:"one_time_keyboard"`
	InputFieldPlaceholder string                  `json:"input_field_placeholder"`
	Selective             bool                    `json:"selective"`
}

type ReplyKeyboardButton struct {
	Text string `json:"text"`
	//RequestUsers    KeyboardButtonRequestUsers `json:"request_users"`
	//RequestChat     KeyboardButtonRequestChat  `json:"request_chat"`
	//RequestContact  bool                       `json:"request_contact"`
	//RequestLocation bool                       `json:"request_location"`
}

type KeyboardButtonRequestUsers struct {
	RequestId       int    `json:"request_id"`
	RequestName     string `json:"request_name"`
	RequestUsername string `json:"request_username"`
	RequestPhoto    string `json:"request_photo"`
}

type KeyboardButtonRequestChat struct {
	RequestId int `json:"request_id"`
}
