package cli

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/kamushadenes/chloe/langchain/memory"
	"os"
)

type ListMessagesCmd struct {
	UserID uint   `arg:"" short:"u" long:"user-id" description:"User ID"`
	Format string `help:"Output format, one of: table, markdown" default:"table"`
}

func (c *ListMessagesCmd) Run(globals *Globals) error {
	u, err := memory.GetUser(globals.Context, c.UserID)
	if err != nil {
		return err
	}

	messages, err := u.ListMessages(globals.Context)
	if err != nil {
		return err
	}

	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleRounded)
	t.Style().Options.SeparateRows = true

	t.AppendHeader(table.Row{"#", "ExternalID", "Interface", "Role", "Content", "Summary", "Token Count", "Moderation Flagged"}, rowConfigAutoMerge)

	for k := range messages {
		message := messages[k]

		moderation := fmt.Sprintf("%t", message.Moderated)

		t.AppendRow([]interface{}{message.ID, message.ExternalID, message.Interface, message.Role, message.Content, message.Summary, message.TokenCount, moderation}, rowConfigAutoMerge)
		t.AppendSeparator()
	}

	switch c.Format {
	case "table":
		t.Render()
	case "markdown":
		t.RenderMarkdown()
	default:
		return fmt.Errorf("unknown format: %s", c.Format)
	}

	return nil
}
