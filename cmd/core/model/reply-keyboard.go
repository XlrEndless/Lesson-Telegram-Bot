package model

type ReplyKeyboardMarkup struct {
	ReplyKeyboardButton [][]ReplyKeyboardButton
	//IsPersistent          bool
	ResizeKeyboard        bool
	OneTimeKeyboard       bool
	InputFieldPlaceholder string
	Selective             bool
}

type ReplyKeyboardButton struct {
	Text string
	//RequestUsers    KeyboardButtonRequestUsers
	//RequestChat     KeyboardButtonRequestChat
	//RequestContact  bool
	//RequestLocation bool
}

type KeyboardButtonRequestUsers struct {
	RequestId       int
	RequestName     string
	RequestUsername string
	RequestPhoto    string
}

type KeyboardButtonRequestChat struct {
	RequestId int
}
