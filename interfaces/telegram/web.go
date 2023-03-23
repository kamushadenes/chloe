package telegram

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kamushadenes/chloe/memory"
	"github.com/rs/zerolog"
	"io"
	"net/http"
)

func getWebPage(ctx context.Context, msg *memory.Message) {
	logger := zerolog.Ctx(ctx)

	resp, err := http.Get(msg.Source.Telegram.Update.Message.Text)
	if err != nil {
		logger.Error().Err(err).Msg("error getting webpage")
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error().Err(err).Msg("error reading webpage")
		return
	}

	err = msg.Save(ctx)

	tmsg := tgbotapi.NewMessage(msg.Source.Telegram.Update.Message.Chat.ID, string(body))
	tmsg.ParseMode = tgbotapi.ModeHTML
	_, _ = msg.Source.Telegram.API.Send(tmsg)
}
