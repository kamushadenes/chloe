package memory

import (
	"context"
	"github.com/kamushadenes/chloe/errors"
	"gorm.io/gorm"
)

func CreateUser(ctx context.Context, firstName, lastName, username string) (*User, error) {
	u := User{
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Mode:      "default",
	}
	u.ID = 0

	if err := db.WithContext(ctx).
		Save(&u).Error; err != nil {
		return nil, errors.Wrap(errors.ErrCreateUser, err)
	}

	return &u, nil
}

func GetUser(ctx context.Context, id uint) (*User, error) {
	var u User

	err := db.WithContext(ctx).
		First(&u, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrap(errors.ErrUserNotFound, err)
		}
		return nil, errors.Wrap(errors.ErrGetUser, err)
	}

	return &u, err
}

func MergeUsersByID(ctx context.Context, ids ...uint) error {
	var users []*User
	for k := range ids {
		u, err := GetUser(ctx, ids[k])
		if err != nil {
			return err
		}
		users = append(users, u)
	}

	return MergeUsers(ctx, users...)
}

func MergeUsers(ctx context.Context, users ...*User) error {
	if len(users) < 2 {
		return nil
	}

	mainUser := users[0]

	for k := range users[1:] {
		user := users[k+1]
		eids, err := user.GetExternalIDs()
		if err != nil {
			return errors.Wrap(errors.ErrGetUser, err)
		}
		for kk := range eids {
			eid := eids[kk]
			if err := mainUser.AddExternalID(ctx, eid.ExternalID, eid.Interface); err != nil {
				return errors.Wrap(errors.ErrUpdateUser, err)
			}
			if err := user.DeleteExternalID(ctx, eid.ExternalID, eid.Interface); err != nil {
				return errors.Wrap(errors.ErrUpdateUser, err)
			}
		}
		if err := BulkChangeMessageOwner(ctx, user, mainUser); err != nil {
			return errors.Wrap(errors.ErrMergeUsers, errors.ErrUpdateMessage, err)
		}
		if err := user.Delete(ctx); err != nil {
			return errors.Wrap(errors.ErrMergeUsers, errors.ErrDeleteUser, err)
		}
	}

	return nil
}

func ListUsers() ([]*User, error) {
	var users []*User
	if err := db.
		Find(&users).Error; err != nil {
		return nil, errors.Wrap(errors.ErrGetUser, err)
	}

	return users, nil
}

func GetUserByExternalID(ctx context.Context, externalID, interf string) (*User, error) {
	var eid ExternalID

	if err := db.WithContext(ctx).
		Where("external_id = ?", externalID).
		Where("interface = ?", interf).
		First(&eid).Error; err != nil {
		return nil, errors.Wrap(errors.ErrGetUser, err)
	}

	var u User

	if err := db.WithContext(ctx).
		First(&u, eid.UserID).Error; err != nil {
		return nil, errors.Wrap(errors.ErrGetUser, err)
	}

	return &u, nil
}

func (u *User) Delete(ctx context.Context) error {
	err := db.WithContext(ctx).
		Delete(u).Error

	if err != nil {
		return errors.Wrap(errors.ErrDeleteUser, err)
	}

	return nil
}

func (u *User) MustGetExternalID(ctx context.Context, interf string) *ExternalID {
	var eid ExternalID

	if tx := db.WithContext(ctx).
		Where("user_id = ?", u.ID).
		Where("interface = ?", interf).
		First(&eid); errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		eid = ExternalID{
			UserID:    u.ID,
			Interface: interf,
		}
		if err := db.WithContext(ctx).
			Create(&eid).Error; err != nil {
			panic(errors.Wrap(errors.ErrCreateUser, err))
		}
	} else if tx.Error != nil {
		panic(errors.Wrap(errors.ErrGetUser, tx.Error))
	}

	return &eid
}

func (u *User) AddExternalID(ctx context.Context, externalID, interf string) error {
	var eid ExternalID

	if tx := db.WithContext(ctx).
		Where("user_id = ?", u.ID).
		Where("interface = ?", interf).
		First(&eid); errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		eid = ExternalID{
			ExternalID: externalID,
			UserID:     u.ID,
			Interface:  interf,
		}
		if err := db.WithContext(ctx).
			Create(&eid).Error; err != nil {
			return errors.Wrap(errors.ErrUpdateUser, err)
		}
	} else if tx.Error != nil {
		return errors.Wrap(errors.ErrGetUser, tx.Error)
	}

	eid.ExternalID = externalID

	err := db.WithContext(ctx).
		Save(&eid).Error

	return errors.Wrap(errors.ErrUpdateUser, err)
}

func (u *User) DeleteExternalID(ctx context.Context, externalID, interf string) error {
	var eid ExternalID

	err := db.WithContext(ctx).
		Where("user_id = ?", u.ID).
		Where("interface = ?", interf).
		Where("external_id = ?", externalID).
		Delete(&eid).Error

	return errors.Wrap(errors.ErrUpdateUser, err)
}

func (u *User) SetMode(ctx context.Context, mode string) error {
	err := db.WithContext(ctx).
		Model(u).
		Update("mode", mode).Error

	return errors.Wrap(errors.ErrUpdateUser, err)
}

func (u *User) Save(ctx context.Context) error {
	err := db.WithContext(ctx).
		Save(u).Error

	return errors.Wrap(errors.ErrSaveUser, err)
}

func (u *User) ListMessages(ctx context.Context) ([]*Message, error) {
	var messages []*Message
	if err := db.WithContext(ctx).
		Where("user_id = ?", u.ID).
		Order("created_at ASC").
		Find(&messages).Error; err != nil {
		return nil, errors.Wrap(errors.ErrLoadMessages, err)
	}

	return messages, nil
}

func (u *User) DeleteMessages(ctx context.Context) error {
	err := db.WithContext(ctx).
		Where("user_id = ?", u.ID).
		Delete(&Message{}).Error

	if err != nil {
		return errors.Wrap(errors.ErrDeleteMessage, err)
	}

	return nil
}

func DeleteAllMessages(ctx context.Context) error {
	err := db.WithContext(ctx).
		Delete(&Message{}).Error

	if err != nil {
		return errors.Wrap(errors.ErrDeleteMessage, err)
	}

	return nil
}

func (u *User) DeleteOldestMessage(ctx context.Context) error {
	var message Message
	if err := db.WithContext(ctx).
		Where("user_id = ?", u.ID).
		Order("created_at").
		First(&message).Error; err != nil {
		return errors.Wrap(errors.ErrLoadMessages, err)
	}

	err := db.WithContext(ctx).
		Delete(&message).Error

	if err != nil {
		return errors.Wrap(errors.ErrDeleteMessage, err)
	}

	return nil
}

func (u *User) GetExternalIDs() ([]*ExternalID, error) {
	var eids []*ExternalID
	if err := db.
		Where("user_id = ?", u.ID).
		Find(&eids).Error; err != nil {
		return nil, errors.Wrap(errors.ErrGetUser, err)
	}

	return eids, nil
}

func (u *User) CreateAPIKey(ctx context.Context) (string, error) {
	apiKey := NewAPIKey(u)

	err := db.WithContext(ctx).Save(apiKey).Error

	if err != nil {
		return "", errors.Wrap(errors.ErrCreateAPIKey, err)
	}

	return apiKey.Key, nil
}
