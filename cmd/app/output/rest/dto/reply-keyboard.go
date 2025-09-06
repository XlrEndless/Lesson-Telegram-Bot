package dto

import "TgBot/cmd/core/model"

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

func MapReplyKeyboardMarkupToModel(input ReplyKeyboardMarkup) model.ReplyKeyboardMarkup {
	replyKeyboardButtonsDto := input.ReplyKeyboardButton
	replyKeyboardButtons := make([][]model.ReplyKeyboardButton, 0)
	for _, arr := range replyKeyboardButtonsDto {
		layout := make([]model.ReplyKeyboardButton, 0)
		for _, button := range arr {
			layout = append(layout, MapReplyKeyboardButtonToModel(button))
		}
		replyKeyboardButtons = append(replyKeyboardButtons, layout)
	}
	return model.ReplyKeyboardMarkup{
		ReplyKeyboardButton:   replyKeyboardButtons,
		ResizeKeyboard:        input.ResizeKeyboard,
		OneTimeKeyboard:       input.OneTimeKeyboard,
		InputFieldPlaceholder: input.InputFieldPlaceholder,
		Selective:             input.Selective,
	}
}

func MapReplyKeyboardMarkupToDto(input model.ReplyKeyboardMarkup) ReplyKeyboardMarkup {
	replyKeyboardButtons := input.ReplyKeyboardButton
	replyKeyboardButtonsDto := make([][]ReplyKeyboardButton, 0)
	for _, arr := range replyKeyboardButtons {
		layout := make([]ReplyKeyboardButton, 0)
		for _, button := range arr {
			layout = append(layout, MapReplyKeyboardButtonToDto(button))
		}
		replyKeyboardButtonsDto = append(replyKeyboardButtonsDto, layout)
	}
	return ReplyKeyboardMarkup{
		ReplyKeyboardButton:   replyKeyboardButtonsDto,
		ResizeKeyboard:        input.ResizeKeyboard,
		OneTimeKeyboard:       input.OneTimeKeyboard,
		InputFieldPlaceholder: input.InputFieldPlaceholder,
		Selective:             input.Selective,
	}
}

func MapReplyKeyboardButtonToModel(input ReplyKeyboardButton) model.ReplyKeyboardButton {
	return model.ReplyKeyboardButton{
		Text: input.Text,
	}
}

func MapReplyKeyboardButtonToDto(input model.ReplyKeyboardButton) ReplyKeyboardButton {
	return ReplyKeyboardButton{
		Text: input.Text,
	}
}
