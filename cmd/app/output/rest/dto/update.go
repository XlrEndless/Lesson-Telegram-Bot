package dto

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
