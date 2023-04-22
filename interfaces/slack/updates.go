package slack

import (
	"context"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/i18n"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/logging"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
	"strings"
)

func HandleUpdates(ctx context.Context, socketMode *socketmode.Client, api *slack.Client, auth *slack.AuthTestResponse) {
	logger := logging.FromContext(ctx)

	for envelope := range socketMode.Events {
		switch envelope.Type {
		case socketmode.EventTypeEventsAPI:
			// Acknowledge the eventPayload first
			socketMode.Ack(*envelope.Request)

			eventPayload, _ := envelope.Data.(slackevents.EventsAPIEvent)

			switch eventPayload.Type {
			case slackevents.CallbackEvent:
				switch event := eventPayload.InnerEvent.Data.(type) {
				case *slackevents.MessageEvent:
					nlogger := logger.With().Str("requestID", event.TimeStamp).Logger()
					ctx = nlogger.WithContext(ctx)

					if !isMentioned(auth, event) {
						continue
					}
					if event.User != "" && event.User != auth.UserID {
						msg := memory.NewMessage(event.TimeStamp, "slack")

						msg.Source.Slack = &memory.SlackMessageSource{
							API:     api,
							Message: event,
						}
						msg.SetContent(strings.TrimSpace(event.Text))
						msg.Role = "user"

						user, err := userFromMessage(ctx, msg)
						if err != nil {
							nlogger.Error().Err(err).Msg("error getting user from message")
							continue
						}
						msg.User = user

						if err := channels.RegisterIncomingMessage(msg); err != nil {
							continue
						}

						if err := complete(ctx, msg); err != nil {
							_ = msg.SendError(err)
							continue
						}
					}
				}
			}
		case socketmode.EventTypeSlashCommand:
			// Acknowledge the eventPayload first
			socketMode.Ack(*envelope.Request)

			nlogger := logger.With().Str("requestID", envelope.Request.EnvelopeID).Logger()
			ctx = nlogger.WithContext(ctx)

			cmd, _ := envelope.Data.(slack.SlashCommand)

			msg := memory.NewMessage(cmd.TriggerID, "slack")

			msg.Source.Slack = &memory.SlackMessageSource{
				API:          api,
				SlashCommand: cmd,
			}
			msg.SetContent(strings.TrimSpace(cmd.Text))
			msg.Role = "user"

			user, err := userFromMessage(ctx, msg)
			if err != nil {
				nlogger.Error().Err(err).Msg("error getting user from message")
				continue
			}
			msg.User = user

			if err := channels.RegisterIncomingMessage(msg); err != nil {
				continue
			}

			switch cmd.Command {
			case "/action":
				if err := action(ctx, msg); err != nil {
					_ = msg.SendError(err)
				}
			case "/generate":
				if err := generate(ctx, msg); err != nil {
					_ = msg.SendError(err)
				}
			case "/tts":
				if err := tts(ctx, msg); err != nil {
					_ = msg.SendError(err)
				}
			case "/forget":
				if err := forgetUser(ctx, msg); err != nil {
					_ = msg.SendError(err)
				}

				if _, _, err := api.PostMessage(cmd.ChannelID, slack.MsgOptionText(i18n.GetForgetText(), false)); err != nil {
					nlogger.Error().Err(err).Msg("error sending message")
				}
			}
		}
	}
}
