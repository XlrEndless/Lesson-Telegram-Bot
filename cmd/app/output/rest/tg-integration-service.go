package rest

import (
	"TgBot/cmd/app/output/rest/dto"
	"TgBot/cmd/core/constant"
	"TgBot/cmd/core/model"
	"TgBot/cmd/core/service"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"io"
	"log/slog"
	"net/http"
	"strconv"
)

type TelegramIntegrationService struct {
	tgConfig   *TgConfig
	httpClient *http.Client
}

func NewTelegramIntegrationService(cfg *TgConfig) service.ITelegramIntegrationService {
	return &TelegramIntegrationService{cfg, http.DefaultClient}
}

func (service *TelegramIntegrationService) GetUpdatesWithOffset(offset int64) ([]model.Update, error) {
	var updates []model.Update
	resp, err := service.httpClient.Get(
		service.tgConfig.BasePath +
			service.tgConfig.BotToken +
			constant.GetUpdatesMethod +
			"?offset=" +
			strconv.FormatInt(offset, 10),
	)
	if err != nil {
		slog.Warn("Something went wrong while GET updates from telegram:", err)
		return updates, err
	}
	slog.Info("GET response Status:", resp.Status)
	slog.Info("GET response Headers:", resp.Header)
	respBody, err := io.ReadAll(resp.Body)
	strBody := string(respBody)
	slog.Info("GET response Body:", strBody)
	update := dto.UpdatesResponse[[]dto.Update]{}
	err = json.Unmarshal(respBody, &update)
	err = resp.Body.Close()
	slog.Info("response Body json:", update)
	if update.Ok {
		err := mapstructure.Decode(update.Result, &updates)
		return updates, err
	} else {
		slog.Warn("Error while pooling updates:", err, update.Description)
		return updates, err
	}
}

func (service *TelegramIntegrationService) SendData(message model.OutputMessage[any]) error {
	var err error
	messageType := message.MessageType
	switch messageType {
	case constant.New:
		{
			slog.Info("Send New Telegram Message")
			var dtoMessage dto.OutputMessage[any]
			err = mapstructure.Decode(message, &dtoMessage)
			var markup any
			markup, err = service.getMarkup(message)
			dtoMessage.ReplyMarkup = markup
			var body []byte
			body, err = json.Marshal(dtoMessage)
			slog.Info("body:", string(body))
			if err == nil {
				_, err = service.checkResponse(func() (*http.Response, error) {
					return service.httpClient.Post(
						service.tgConfig.BasePath+service.tgConfig.BotToken+constant.SendMessageMethod,
						constant.ContentType,
						bytes.NewBuffer(body))
				}, constant.SendMessageMethod)
			}
		}
	case constant.Change:
		{
			slog.Info("Send Edit Inline Telegram Message")
			_, err = service.checkResponse(func() (*http.Response, error) {
				callbackData := message.CallbackData
				return service.httpClient.Post(service.tgConfig.BasePath+
					service.tgConfig.BotToken+
					constant.AnswerCallBackMethod+
					"?callback_query_id="+
					callbackData.CallbackId,
					constant.ContentType, nil)
			}, constant.AnswerCallBackMethod)
			var dtoMessage dto.OutputMessage[any]
			err = mapstructure.Decode(message, &dtoMessage)
			var markup any
			markup, err = service.getMarkup(message)
			dtoMessage.ReplyMarkup = markup
			editMessage := dto.EditMessageReplyMarkup{
				ChatId:      message.ChatId,
				MessageId:   message.MessageId,
				ReplyMarkup: markup,
			}
			var body []byte
			body, err = json.Marshal(editMessage)
			slog.Info("body:", string(body))
			if err == nil {
				_, err = service.checkResponse(func() (*http.Response, error) {
					return service.httpClient.Post(
						service.tgConfig.BasePath+service.tgConfig.BotToken+constant.EditInlineMessageMethod,
						constant.ContentType,
						bytes.NewBuffer(body))
				}, constant.EditInlineMessageMethod)
			}
		}
	default:
		{
			slog.Warn("Unknown message type:", messageType)
			err = errors.New("unknown message type")
		}
	}

	return err
}

func (service *TelegramIntegrationService) getMarkup(message model.OutputMessage[any]) (any, error) {
	var err error
	var markup any
	switch message.MarkupType {
	case constant.Reply:
		{
			markup = message.ReplyMarkup
			replyKeyboardMarkup := markup.(model.ReplyKeyboardMarkup)
			var replyKeyboardMarkupDto dto.ReplyKeyboardMarkup
			err = mapstructure.Decode(replyKeyboardMarkup, &replyKeyboardMarkupDto)
			return replyKeyboardMarkupDto, err
		}
	case constant.Inline:
		{
			markup = message.ReplyMarkup
			inlineKeyboardMarkup := markup.(model.InlineKeyboardMarkup)
			var inlineKeyboardMarkupDto dto.InlineKeyboardMarkup
			err = mapstructure.Decode(inlineKeyboardMarkup, &inlineKeyboardMarkupDto)
			return inlineKeyboardMarkupDto, err
		}
	default:
		return markup, err
	}
}

func (service *TelegramIntegrationService) checkResponse(fn func() (*http.Response, error), method string) (*http.Response, error) {
	slog.Info(fmt.Sprintf("Call %s method", method))
	response, err := fn()
	if err != nil {
		slog.Warn(fmt.Sprintf("Something went wrong while call %s:", method), err)
	} else {
		slog.Info("POST response Status:", response.Status)
		slog.Info("POST response Headers:", response.Header)
		if response.StatusCode != http.StatusOK {
			slog.Warn(fmt.Sprintf("Something went wrong while call %s:", method), response.StatusCode)
			err = errors.New(fmt.Sprintf("Something went wrong while call %s", method))
		} else {
			slog.Warn(fmt.Sprintf("Successfuly get response from %s:", method))
		}
	}
	return response, err
}
