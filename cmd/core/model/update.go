package model

import (
	"TgBot/cmd/core/constant"
	"strings"
)

var commandsMap = map[string]bool{
	constant.NextPageLabel:        true,
	constant.ChooseDateLabel:      true,
	constant.ShowMyLessonsLabel:   true,
	constant.PreviousCommandLabel: true,
	constant.RegToLessonLabel:     true,
}

type Update struct {
	ID                 int
	Message            InputMessage
	CallbackQuery      CallbackQuery
	InlineQuery        InlineQuery
	ChosenInlineResult ChosenInlineResult
}

type InlineQuery struct {
	Id       string
	From     User
	Query    string
	Offset   string
	ChatType string
}

type ChosenInlineResult struct {
	ResultID        string
	From            User
	InlineMessageID string
	Query           string
}

func (update *Update) IsCommand() bool {
	messageText := update.Message.Text
	return strings.HasPrefix(messageText, "/")
}

func (update *Update) IsQuery() bool {
	callbackData := update.CallbackQuery.Data
	return strings.HasPrefix(callbackData, "/")
}

func (update *Update) IsNotPrefixedCommand() bool {
	messageText := update.Message.Text
	return commandsMap[messageText]
}
