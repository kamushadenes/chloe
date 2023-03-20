package users

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/memory"
)

type ExternalID struct {
	ID        string `json:"id"`
	Interface string `json:"interface"`
}

type User struct {
	ID         string      `json:"id"`
	ExternalID *ExternalID `json:"external_id"`
	FirstName  string      `json:"first_name"`
	LastName   string      `json:"last_name"`
	Username   string      `json:"username"`
}

func CreateUser(ctx context.Context, firstName, lastName, username string) (*User, error) {
	id := generateId()
	err := memory.CreateUser(ctx, id, firstName, lastName, username)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
	}, nil
}

func GetUser(ctx context.Context, id string) (*User, error) {
	firstName, lastName, username, err := memory.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
	}, nil
}

func GetUserOrSetByExternalId(ctx context.Context, externalId, interf, firstName, lastName, username string) (*User, error) {
	var user *User
	var err error

	user, err = GetUserByExternalId(ctx, externalId, interf)
	if err != nil {
		user, err = CreateUser(ctx, firstName, lastName, username)
		if err != nil {
			return nil, err
		}
		if err = AddUserExternalId(ctx, user.ID, externalId, interf); err != nil {
			return nil, err
		}
		user, err = GetUserByExternalId(ctx, externalId, interf)
	}

	return user, err
}

func GetUserByExternalId(ctx context.Context, externalId, interf string) (*User, error) {
	id, err := memory.FindUserByExternalId(ctx, externalId, interf)
	if err != nil {
		return nil, err
	}

	user, err := GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	user.ExternalID = &ExternalID{
		ID:        externalId,
		Interface: interf,
	}

	return user, nil
}

func generateId() string {
	// random uuid
	id, _ := uuid.Must(uuid.NewV4()).MarshalText()
	return string(id)
}

func AddUserExternalId(ctx context.Context, userId, externalId, interf string) error {
	return memory.AddUserExternalId(ctx, userId, externalId, interf)
}
