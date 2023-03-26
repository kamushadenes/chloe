package memory

import (
	"context"
)

func NewUser(ctx context.Context, firstName, lastName, username string) (*User, error) {
	u := User{
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Mode:      "default",
	}
	u.ID = 0

	if err := db.WithContext(ctx).Save(&u).Error; err != nil {
		return nil, err
	}

	return &u, nil
}

func GetUser(ctx context.Context, id string) (*User, error) {
	var u User

	err := db.WithContext(ctx).First(&u, id).Error

	return &u, err
}

func GetUserByExternalID(ctx context.Context, externalID, interf string) (*User, error) {
	var eid ExternalID

	if err := db.WithContext(ctx).First(&eid, "external_id = ? AND interface = ?", externalID, interf).Error; err != nil {
		return nil, err
	}

	var u User

	if err := db.WithContext(ctx).First(&u, eid.UserID).Error; err != nil {
		return nil, err
	}

	return &u, nil
}
