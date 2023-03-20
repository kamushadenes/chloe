package memory

import (
	"context"
)

func CreateUser(ctx context.Context, userId, firstName, lastName, username string) error {
	_, err := db.ExecContext(ctx,
		"INSERT INTO users (id, first_name, last_name, username) VALUES (?, ?, ?, ?) "+
			"ON CONFLICT(id) DO UPDATE SET first_name = ?, last_name = ?, username = ?",
		userId, firstName, lastName, username, firstName, lastName, username)

	return err
}

func AddUserExternalId(ctx context.Context, userId, externalId, interf string) error {
	_, err := db.ExecContext(ctx,
		"INSERT INTO user_external_ids (user_id, external_id, interface) VALUES (?, ?, ?) "+
			"ON CONFLICT (external_id, interface) DO UPDATE SET user_id = ?",
		userId, externalId, interf, userId)

	return err
}

func GetUser(ctx context.Context, userId string) (string, string, string, error) {
	var firstName, lastName, username string
	err := db.QueryRowContext(ctx,
		"SELECT first_name, last_name, username FROM users WHERE id = ?", userId).
		Scan(&firstName, &lastName, &username)

	return firstName, lastName, username, err
}

func FindUserByExternalId(ctx context.Context, externalId, interf string) (string, error) {
	var id string
	if err := db.QueryRowContext(ctx,
		"SELECT user_id FROM user_external_ids WHERE external_id = ? AND interface = ?", externalId, interf).Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func GetUserMode(ctx context.Context, userId string) (string, error) {
	var mode string
	if err := db.QueryRowContext(ctx,
		"SELECT mode FROM modes WHERE user_id = ?", userId).Scan(&mode); err == nil {
		return mode, nil
	}

	_, err := db.Exec("INSERT INTO modes (user_id, mode) VALUES (?, ?)", userId, "default")

	return "default", err
}

func SetUserMode(ctx context.Context, userId, mode string) error {
	_, err := db.ExecContext(ctx,
		"INSERT INTO modes (user_id, mode) VALUES (?, ?) "+
			"ON CONFLICT(user_id) DO UPDATE SET mode = ?", userId, mode, mode)

	return err
}
