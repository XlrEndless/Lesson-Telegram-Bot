package dto

import "TgBot/cmd/core/model"

type UpdatesResponse[T any] struct {
	Ok          bool   `json:"ok"`
	Description string `json:"description"`
	Result      T      `json:"result"`
}

type Update struct {
	ID                 int                `json:"update_id"`
	Message            InputMessage       `json:"message"`
	CallbackQuery      CallbackQuery      `json:"callback_query"`
	InlineQuery        InlineQuery        `json:"inline_query"`
	ChosenInlineResult ChosenInlineResult `json:"chosen_inline_result"`
}

type InlineQuery struct {
	Id       string `json:"id"`
	From     User   `json:"from"`
	Query    string `json:"query"`
	Offset   string `json:"offset"`
	ChatType string `json:"chat_type"`
}

type ChosenInlineResult struct {
	ResultID        string `json:"result_id"`
	From            User   `json:"from"`
	InlineMessageID string `json:"inline_message_id"`
	Query           string `json:"query"`
}

func MapUpdatesResponseToModel(input UpdatesResponse[any]) model.UpdatesResponse[any] {
	return model.UpdatesResponse[any]{
		Ok:          input.Ok,
		Description: input.Description,
		Result:      input.Result,
	}
}

func MapUpdateToModel(input Update) model.Update {
	return model.Update{
		ID:                 input.ID,
		Message:            MapInputMessage(input.Message),
		CallbackQuery:      MapCallbackQuery(input.CallbackQuery),
		InlineQuery:        MapInlineQueryToModel(input.InlineQuery),
		ChosenInlineResult: MapChosenInlineResultToModel(input.ChosenInlineResult),
	}
}

func MapInlineQueryToModel(input InlineQuery) model.InlineQuery {
	return model.InlineQuery{
		Id:       input.Id,
		From:     MapUserToModel(input.From),
		Query:    input.Query,
		Offset:   input.Offset,
		ChatType: input.ChatType,
	}
}

func MapChosenInlineResultToModel(input ChosenInlineResult) model.ChosenInlineResult {
	return model.ChosenInlineResult{
		ResultID:        input.ResultID,
		From:            MapUserToModel(input.From),
		InlineMessageID: input.InlineMessageID,
		Query:           input.Query,
	}
}
