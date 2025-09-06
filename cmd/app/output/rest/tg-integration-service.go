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
	resp, err := service.logResponse(func() (*http.Response, error) {
		return service.httpClient.Get(
			service.tgConfig.BasePath +
				service.tgConfig.BotToken +
				constant.GetUpdatesMethod +
				"?offset=" +
				strconv.FormatInt(offset, 10),
		)
	}, constant.GetUpdatesMethod)
	if err != nil {
		slog.Warn("Something went wrong while GET updates from telegram:", err)
		return updates, err
	}
	respBody, err := io.ReadAll(resp.Body)
	strBody := string(respBody)
	slog.Info("GET response Body:", strBody)
	updateResponseDto := dto.UpdatesResponse[[]dto.Update]{}
	err = json.Unmarshal(respBody, &updateResponseDto)
	err = resp.Body.Close()
	slog.Info("response Body json:", updateResponseDto)
	if updateResponseDto.Ok {
		updates := make([]model.Update, 0)
		for _, updateDto := range updateResponseDto.Result {
			updates = append(updates, dto.MapUpdateToModel(updateDto))
		}
		return updates, err
	} else {
		slog.Warn("Error while pooling updates:", err, updateResponseDto.Description)
		return updates, err
	}
}

func (service *TelegramIntegrationService) SendData(message model.OutputMessage[any]) error {
	err := service.answerUserIfNeeded(message)
	err = service.deleteMessageIfNeeded(message)
	messageType := message.MessageType
	switch messageType {
	case constant.New:
		{
			slog.Info("Send New Telegram Message")
			dtoMessage := dto.MapOutputMessageDto(message)
			var body []byte
			body, err = json.Marshal(dtoMessage)
			slog.Info("body:", string(body))
			if err == nil {
				_, err = service.logResponse(func() (*http.Response, error) {
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
			editMessage := dto.EditMessageReplyMarkup{
				ChatId:      message.ChatId,
				MessageId:   message.MessageId,
				ReplyMarkup: dto.MapInlineKeyboardMarkupToDto(message.ReplyMarkup.(model.InlineKeyboardMarkup)),
			}
			var body []byte
			body, err = json.Marshal(editMessage)
			slog.Info("body:", string(body))
			if err == nil {
				_, err = service.logResponse(func() (*http.Response, error) {
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

func (service *TelegramIntegrationService) logResponse(fn func() (*http.Response, error), method string) (*http.Response, error) {
	slog.Info(fmt.Sprintf("Call %s method", method))
	response, err := fn()
	if err != nil {
		slog.Warn(fmt.Sprintf("Something went wrong while call %s:", method), err)
	} else {
		slog.Info(method+" response Status:", response.Status)
		slog.Info(method+" response Headers:", response.Header)
		if response.StatusCode != http.StatusOK {
			slog.Warn(fmt.Sprintf("Something went wrong while call %s:", method), response.StatusCode)
			err = errors.New(fmt.Sprintf("Something went wrong while call %s", method))
		} else {
			slog.Warn(fmt.Sprintf("Successfuly get response from %s:", method))
		}
	}
	return response, err
}

func (service *TelegramIntegrationService) deleteMessageIfNeeded(message model.OutputMessage[any]) error {
	if message.ShouldDelete {
		slog.Info("Delete Telegram Message")
		_, err := service.logResponse(func() (*http.Response, error) {
			return service.httpClient.Post(service.tgConfig.BasePath+
				service.tgConfig.BotToken+
				constant.DeleteMessageMethod+
				"?chat_id="+
				strconv.Itoa(message.ChatId)+
				"&message_id="+
				strconv.Itoa(message.MessageId),
				constant.ContentType, nil)
		}, constant.DeleteMessageMethod)
		if err != nil {
			slog.Warn("Can't delete message:", err)
		}
		return err
	} else {
		return nil
	}
}

func (service *TelegramIntegrationService) answerUserIfNeeded(message model.OutputMessage[any]) error {
	if message.NeedAnswer {
		slog.Info("Send Answer Telegram Message")
		_, err := service.logResponse(func() (*http.Response, error) {
			callbackData := message.CallbackData
			return service.httpClient.Post(service.tgConfig.BasePath+
				service.tgConfig.BotToken+
				constant.AnswerCallBackMethod+
				"?callback_query_id="+
				callbackData.CallbackId,
				constant.ContentType, nil)
		}, constant.AnswerCallBackMethod)
		if err != nil {
			slog.Warn("Can't send Answer to user:", err)
		}
		return err
	} else {
		return nil
	}
}
