package cli

import (
	"context"
	"time"

	"github.com/briandowns/spinner"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/colors"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/flags"
	"github.com/kamushadenes/chloe/langchain/chat_models"
	"github.com/kamushadenes/chloe/langchain/chat_models/messages"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs/writer_structs"
)

func Complete(ctx context.Context, text string, writer writer_structs.ChloeWriter) error {
	s := spinner.New(spinner.CharSets[40], 100*time.Millisecond)

	if flags.InteractiveCLI {
		s.Prefix = colors.BoldCyan("Chloe: ")
		s.FinalMSG = s.Prefix
		s.Start()
	}

	msg := memory.NewMessage(uuid.Must(uuid.NewV4()).String(), "cli")
	msg.Role = "user"
	msg.User = user
	msg.Source = &memory.MessageSource{
		CLI: &memory.CLIMessageSource{},
	}

	msg.SetContent(text)

	if err := msg.Save(ctx); err != nil {
		return err
	}

	chat := chat_models.NewChatWithDefaultModel(config.Chat.Provider, msg.User)

	if flags.InteractiveCLI {
		writer.SetPreWriteCallback(func() {
			s.Stop()
		})
	}

	_, err := chat.ChatStreamWithContext(ctx, writer, msg, messages.UserMessage(text))

	return err
}
