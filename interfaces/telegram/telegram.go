package telegram

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/logging"
	"github.com/rs/zerolog"
)

func newBot(ctx context.Context, token string) (*tgbotapi.BotAPI, error) {
	logger := zerolog.Ctx(ctx)

	logger.Info().Msg("creating telegram bot")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, errors.Wrap(errors.ErrCreateTelegramBot, err)
	}

	logger.Info().Str("account", bot.Self.UserName).Msg("telegram bot created")

	return bot, err
}

func Start(ctx context.Context) {
	logger := logging.GetLogger().With().Str("interface", "telegram").Logger()
	ctx = logger.WithContext(ctx)

	if config.Telegram.Token == "" {
		logger.Warn().Msg("token not configured, telegram interface disabled")
		return
	}

	logger.Info().Msg("starting telegram interface")

	bot, err := newBot(ctx, config.Telegram.Token)
	if err != nil {
		logger.Panic().Err(err).Msg("error in telegram interface")
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for {
		select {
		case <-ctx.Done():
			logger.Warn().Err(ctx.Err()).Msg("closing telegram interface")
			return
		case update := <-updates:
			go handleUpdates(ctx, bot, update)
		}
	}
}
