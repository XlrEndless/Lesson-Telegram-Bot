package impl

import (
	"TgBot/cmd/core/service"
	"TgBot/cmd/core/tools"
	"context"
	"log/slog"
	"time"
)

type TelegramMessageConsumer struct {
	telegramService    service.ITelegramService
	pollIntervalMillis int64
}

func NewTelegramMessageConsumer(telegramService service.ITelegramService, cfg *tools.CommonConfig) *TelegramMessageConsumer {
	return &TelegramMessageConsumer{telegramService: telegramService, pollIntervalMillis: cfg.PollIntervalMillis}
}

func (consumer *TelegramMessageConsumer) ConsumeMessages(ctx context.Context) error {
	for {
		slog.Warn("Polling telegram updates...")
		updates, err := consumer.telegramService.GetUpdates(ctx)
		if err != nil {
			slog.Warn("Error getting updates: %v", err)
		}
		slog.Warn("Updates size:", len(updates))
		for _, update := range updates {
			err = consumer.telegramService.HandleTelegramUpdate(ctx, update)
			if err != nil {
				slog.Warn("Something went wrong when handle telegram update: %v", err)
				break
			}
		}
		if err == nil {
			slog.Warn("Successfully handled telegram updates!")
		}
		slog.Warn("Wait another poll...")
		time.Sleep(time.Duration(consumer.pollIntervalMillis) * time.Millisecond)
	}
}
