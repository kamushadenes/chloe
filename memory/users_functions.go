package memory

import (
	"context"
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
		return nil, err
	}

	return &u, nil
}

func GetUser(ctx context.Context, id uint) (*User, error) {
	var u User

	err := db.WithContext(ctx).
		First(&u, id).Error

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
			return err
		}
		for kk := range eids {
			eid := eids[kk]
			if err := mainUser.AddExternalID(ctx, eid.ExternalID, eid.Interface); err != nil {
				return err
			}
			if err := user.DeleteExternalID(ctx, eid.ExternalID, eid.Interface); err != nil {
				return err
			}
		}
		if err := BulkChangeMessageOwner(ctx, user, mainUser); err != nil {
			return err
		}
		if err := user.Delete(ctx); err != nil {
			return err
		}
	}

	return nil
}

func ListUsers() ([]*User, error) {
	var users []*User
	if err := db.
		Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByExternalID(ctx context.Context, externalID, interf string) (*User, error) {
	var eid ExternalID

	if err := db.WithContext(ctx).
		Where("external_id = ?", externalID).
		Where("interface = ?", interf).
		First(&eid).Error; err != nil {
		return nil, err
	}

	var u User

	if err := db.WithContext(ctx).
		First(&u, eid.UserID).Error; err != nil {
		return nil, err
	}

	return &u, nil
}

func (u *User) Delete(ctx context.Context) error {
	return db.WithContext(ctx).
		Delete(u).Error
}

func (u *User) MustGetExternalID(ctx context.Context, interf string) *ExternalID {
	var eid ExternalID

	if tx := db.WithContext(ctx).
		Where("user_id = ?", u.ID).
		Where("interface = ?", interf).
		First(&eid); tx.Error == gorm.ErrRecordNotFound {
		eid = ExternalID{
			UserID:    u.ID,
			Interface: interf,
		}
		if err := db.WithContext(ctx).
			Create(&eid).Error; err != nil {
			panic(err)
		}
	} else if tx.Error != nil {
		panic(tx.Error)
	}

	return &eid
}

func (u *User) AddExternalID(ctx context.Context, externalID, interf string) error {
	var eid ExternalID

	if tx := db.WithContext(ctx).
		Where("user_id = ?", u.ID).
		Where("interface = ?", interf).
		First(&eid); tx.Error == gorm.ErrRecordNotFound {
		eid = ExternalID{
			ExternalID: externalID,
			UserID:     u.ID,
			Interface:  interf,
		}
		if err := db.WithContext(ctx).
			Create(&eid).Error; err != nil {
			return err
		}
	} else if tx.Error != nil {
		return tx.Error
	}

	eid.ExternalID = externalID

	return db.WithContext(ctx).
		Save(&eid).Error
}

func (u *User) DeleteExternalID(ctx context.Context, externalID, interf string) error {
	var eid ExternalID

	return db.WithContext(ctx).
		Where("user_id = ?", u.ID).
		Where("interface = ?", interf).
		Where("external_id = ?", externalID).
		Delete(&eid).Error
}

func (u *User) SetMode(ctx context.Context, mode string) error {
	return db.WithContext(ctx).
		Model(u).
		Update("mode", mode).Error
}

func (u *User) Save(ctx context.Context) error {
	return db.WithContext(ctx).
		Save(u).Error
}

func (u *User) ListMessages(ctx context.Context) ([]*Message, error) {
	var messages []*Message
	if err := db.WithContext(ctx).
		Where("user_id = ?", u.ID).
		Order("created_at ASC").
		Find(&messages).Error; err != nil {
		return nil, err
	}

	return messages, nil
}

func (u *User) DeleteMessages(ctx context.Context) error {
	return db.WithContext(ctx).
		Where("user_id = ?", u.ID).
		Delete(&Message{}).Error
}

func DeleteAllMessages(ctx context.Context) error {
	return db.WithContext(ctx).
		Delete(&Message{}).Error
}

func (u *User) DeleteOldestMessage(ctx context.Context) error {
	var message Message
	if err := db.WithContext(ctx).
		Where("user_id = ?", u.ID).
		Order("created_at").
		First(&message).Error; err != nil {
		return err
	}

	return db.WithContext(ctx).
		Delete(&message).Error
}

func (u *User) GetExternalIDs() ([]*ExternalID, error) {
	var eids []*ExternalID
	if err := db.
		Where("user_id = ?", u.ID).
		Find(&eids).Error; err != nil {
		return nil, err
	}

	return eids, nil
}

func (u *User) CreateAPIKey(ctx context.Context) (string, error) {
	apiKey := NewAPIKey(u)

	return apiKey.Key, db.WithContext(ctx).Save(apiKey).Error
}
