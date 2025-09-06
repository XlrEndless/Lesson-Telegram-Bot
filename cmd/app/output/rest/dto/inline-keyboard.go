package dto

import "TgBot/cmd/core/model"

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}

func MapInlineKeyboardMarkupToModel(input InlineKeyboardMarkup) model.InlineKeyboardMarkup {
	inlineKeyboard := make([][]model.InlineKeyboardButton, 0)
	for _, arr := range input.InlineKeyboard {
		layout := make([]model.InlineKeyboardButton, 0)
		for _, button := range arr {
			layout = append(layout, MapInlineKeyboardButtonToModel(button))
		}
		inlineKeyboard = append(inlineKeyboard, layout)
	}
	return model.InlineKeyboardMarkup{
		InlineKeyboard: inlineKeyboard,
	}
}

func MapInlineKeyboardMarkupToDto(input model.InlineKeyboardMarkup) InlineKeyboardMarkup {
	inlineKeyboardDto := make([][]InlineKeyboardButton, 0)
	for _, arr := range input.InlineKeyboard {
		layout := make([]InlineKeyboardButton, 0)
		for _, button := range arr {
			layout = append(layout, MapInlineKeyboardButtonToDto(button))
		}
		inlineKeyboardDto = append(inlineKeyboardDto, layout)
	}
	return InlineKeyboardMarkup{
		InlineKeyboard: inlineKeyboardDto,
	}
}

func MapInlineKeyboardButtonToModel(input InlineKeyboardButton) model.InlineKeyboardButton {
	return model.InlineKeyboardButton{
		Text:         input.Text,
		CallbackData: input.CallbackData,
	}
}

func MapInlineKeyboardButtonToDto(input model.InlineKeyboardButton) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text:         input.Text,
		CallbackData: input.CallbackData,
	}
}
