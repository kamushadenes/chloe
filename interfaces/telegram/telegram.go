package telegram

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog"
	"os"
)

func newBot(ctx context.Context, token string) (*tgbotapi.BotAPI, error) {
	logger := zerolog.Ctx(ctx)

	logger.Info().Msg("creating telegram bot")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("error creating telegram bot: %w", err)
	}

	logger.Info().Str("account", bot.Self.UserName).Msg("telegram bot created")

	return bot, err
}

func Start(ctx context.Context) {
	logger := zerolog.Ctx(ctx).With().Str("interface", "telegram").Logger()
	ctx = logger.WithContext(ctx)

	logger.Info().Msg("starting telegram interface")

	bot, err := newBot(ctx, os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		logger.Panic().Err(err).Msg("error in telegram interface")
	}

	//bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for {
		select {
		case <-ctx.Done():
			logger.Panic().Err(ctx.Err()).Msg("error in telegram interface")
		case update := <-updates:
			go handleUpdates(ctx, bot, update)
		}
	}
}
