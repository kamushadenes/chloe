package slack

import (
	"context"

	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/logging"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func newBot(ctx context.Context, token string, appLevelToken string) (*slack.Client, error) {
	api := slack.New(
		token,
		slack.OptionAppLevelToken(appLevelToken),
	)

	return api, nil
}

func Start(ctx context.Context) {
	logger := logging.GetLogger().With().Str("interface", "slack").Logger()
	ctx = logger.WithContext(ctx)

	if config.Slack.Token == "" {
		logger.Warn().Msg("token not configured, slack interface disabled")
		return
	}

	api, _ := newBot(ctx, config.Slack.Token, config.Slack.AppLevelToken)

	socketMode := socketmode.New(api)
	auth, err := api.AuthTest()
	if err != nil {
		logger.Error().Msg("invalid credentials")
		return
	}

	logger.Info().Str("account", auth.User).Msg("slack bot created")

	go HandleUpdates(ctx, socketMode, api, auth)

	if err := socketMode.Run(); err != nil {
		logger.Error().Err(err).Msg("error in slack interface")
		return
	}
}
