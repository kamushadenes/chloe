package slack

import (
	"context"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/i18n"
	"github.com/kamushadenes/chloe/memory"
	"github.com/rs/zerolog"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
	"strings"
)

func HandleUpdates(ctx context.Context, socketMode *socketmode.Client, api *slack.Client, auth *slack.AuthTestResponse) {
	logger := zerolog.Ctx(ctx)

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
							logger.Error().Err(err).Msg("error getting user from message")
							continue
						}
						msg.User = user

						channels.IncomingMessagesCh <- msg
						if err := <-msg.ErrorCh; err != nil {
							logger.Error().Msg("error saving message")
							continue
						}

						complete(ctx, msg)
					}
				}
			}
		case socketmode.EventTypeSlashCommand:
			// Acknowledge the eventPayload first
			socketMode.Ack(*envelope.Request)

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
				logger.Error().Err(err).Msg("error getting user from message")
				continue
			}
			msg.User = user

			channels.IncomingMessagesCh <- msg
			if err := <-msg.ErrorCh; err != nil {
				logger.Error().Msg("error saving message")
				continue
			}

			switch cmd.Command {
			case "/action":
				action(ctx, msg)
			case "/generate":
				generate(ctx, msg)
			case "/tts":
				tts(ctx, msg)
			case "/forget":
				if err := forgetUser(ctx, msg); err != nil {
					logger.Error().Err(err).Msg("error forgetting user")
				}

				if _, _, err := api.PostMessage(cmd.ChannelID, slack.MsgOptionText(i18n.GetForgetText(), false)); err != nil {
					logger.Error().Err(err).Msg("error sending message")
				}
			}
		}
	}
}
