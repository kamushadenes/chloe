package memory

import (
	"context"
	"database/sql"
	"github.com/gofrs/uuid"
)

func SaveMessage(ctx context.Context, userId, role, content string, chainOfThought string) error {
	id := uuid.Must(uuid.NewV4())
	_, err := db.ExecContext(ctx,
		"INSERT INTO messages (id, user_id, role, content, chain_of_thought) VALUES (?, ?, ?, ?, ?)",
		id, userId, role, content, chainOfThought)

	return err
}

func LoadMessages(ctx context.Context, userId string) ([][]string, error) {
	rows, err := db.QueryContext(ctx,
		"SELECT id, role, content, summary, chain_of_thought FROM messages WHERE user_id = ? "+
			"ORDER BY created_at", userId)
	if err != nil {
		return nil, err
	}

	var messages [][]string

	for rows.Next() {
		var id, role, content, summary, chainOfThought sql.NullString
		err := rows.Scan(&id, &role, &content, &summary, &chainOfThought)
		if err != nil {
			return nil, err
		}

		messages = append(messages, []string{id.String, role.String, content.String, summary.String, chainOfThought.String})
	}

	return messages, nil
}

func DeleteMessages(ctx context.Context, userId string) error {
	_, err := db.ExecContext(ctx,
		"DELETE FROM messages WHERE user_id = ?", userId)

	return err
}

func SetMessageSummary(ctx context.Context, id string, summary string) error {
	_, err := db.ExecContext(ctx,
		"UPDATE messages SET summary = ? WHERE id = ?", summary, id)

	return err
}

func LoadNonSummarizedMessages(ctx context.Context) ([][]string, error) {
	rows, err := db.QueryContext(ctx,
		"SELECT id, content FROM messages WHERE summary IS NULL AND content IS NOT NULL")
	if err != nil {
		return nil, err
	}

	var messages [][]string

	for rows.Next() {
		var id, content string
		err := rows.Scan(&id, &content)
		if err != nil {
			return nil, err
		}

		messages = append(messages, []string{id, content})
	}

	return messages, nil
}
