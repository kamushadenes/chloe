package discord

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/logging"
	"github.com/rs/zerolog"
)

func newBot(ctx context.Context, token string) (*discordgo.Session, error) {
	logger := zerolog.Ctx(ctx)

	logger.Info().Msg("creating discord bot")

	bot, err := discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		return nil, fmt.Errorf("error creating discord bot: %w", err)
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

	bot.Identify.Intents = discordgo.MakeIntent(
		discordgo.IntentsGuildMessages |
			discordgo.IntentsDirectMessages |
			discordgo.IntentsMessageContent)

	bot.AddHandler(handleMessageCreate)
	bot.AddHandler(handleCommandInteraction)

	err = bot.Open()
	if err != nil {
		logger.Error().Err(err).Msg("error in discord interface")
		return
	}

	if err := registerCommands(bot); err != nil {
		logger.Error().Err(err).Msg("error registering commands")
		return
	}

	for {
		select {
		case <-ctx.Done():
			logger.Warn().Err(ctx.Err()).Msg("closing discord interface")
			_ = bot.Close()
			return
		}
	}
}
