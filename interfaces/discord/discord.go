package discord

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/utils"
	"github.com/rs/zerolog"
	"time"
)

var handlers = []interface{}{
	handleMessageCreate,
	handleCommandInteraction,
}

var intents = []discordgo.Intent{
	discordgo.IntentsGuildMessages,
	discordgo.IntentsDirectMessages,
	discordgo.IntentsMessageContent,
}

func newBot(ctx context.Context, token string) (*discordgo.Session, error) {
	logger := zerolog.Ctx(ctx)

	logger.Info().Msg("creating discord bot")

	bot, err := discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		return nil, errors.Wrap(errors.ErrCreateDiscordBot, err)
	}

	bot.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		logger.Info().Str("account", fmt.Sprintf("%s#%s", bot.State.User.Username, bot.State.User.Discriminator)).Msg("discord bot created")
	})

	return bot, err
}

func Start(ctx context.Context) {
	logger := logging.GetLogger().With().Str("interface", "discord").Logger()
	ctx = logger.WithContext(ctx)

	if config.Discord.Token == "" {
		logger.Warn().Msg("token not configured, discord interface disabled")
		return
	}

	logger.Info().Msg("starting discord interface")

	bot, err := newBot(ctx, config.Discord.Token)
	if err != nil {
		logger.Error().Err(err).Msg("error in discord interface")
		return
	}

	for k := range intents {
		bot.Identify.Intents |= intents[k]
	}

	for k := range handlers {
		bot.AddHandler(handlers[k])
	}

	err = bot.Open()
	if err != nil {
		logger.Error().Err(err).Msg("error in discord interface")
		return
	}

	if err := registerCommands(bot); err != nil {
		logger.Error().Err(err).Msg("error registering commands")
		return
	}

	ticker := utils.TickerOrDefault(config.Discord.RandomStatusUpdateInterval, 10*time.Minute)

	for {
		select {
		case <-ctx.Done():
			logger.Warn().Err(ctx.Err()).Msg("closing discord interface")
			_ = bot.Close()
			return
		case <-ticker.C:
			if config.Discord.RandomStatusUpdateInterval > 0 {
				if err := updateStatus(bot); err != nil {
					logger.Error().Err(err).Msg("error updating status")
				}
			}
		}
	}
}
