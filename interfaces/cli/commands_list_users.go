package cli

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/kamushadenes/chloe/langchain/memory"
	"os"
)

type ListUsersCmd struct {
	Format string `help:"Output format, one of: table, markdown" default:"table"`
}

func (c *ListUsersCmd) Run(globals *Globals) error {
	users, err := memory.ListUsers()
	if err != nil {
		return err
	}

	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleRounded)
	t.Style().Options.SeparateRows = true

	t.AppendHeader(table.Row{"#", "First Name", "Last Name", "Username", "Mode", "External IDs", "External IDs"}, rowConfigAutoMerge)
	t.AppendHeader(table.Row{"", "", "", "", "", "Interface", "ID"})

	t.SetColumnConfigs([]table.ColumnConfig{
		{
			Name:      "#",
			AutoMerge: true,
		},
		{
			Name:      "First Name",
			AutoMerge: true,
		},
		{
			Name:      "Last Name",
			AutoMerge: true,
		},
		{
			Name:      "Username",
			AutoMerge: true,
		},
		{
			Name:      "Mode",
			AutoMerge: true,
		},
	})

	for k := range users {
		u := users[k]

		eids, err := u.GetExternalIDs()
		if err != nil {
			return err
		}

		if len(eids) == 0 {
			t.AppendRow([]interface{}{u.ID, u.FirstName, u.LastName, u.Username, u.Mode, "", ""})
		} else {
			for kk := range eids {
				t.AppendRow([]interface{}{u.ID, u.FirstName, u.LastName, u.Username, u.Mode, eids[kk].Interface, eids[kk].ExternalID})
			}
		}
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
