package memory

import (
	"context"
	"gorm.io/gorm"
)

type ExternalID struct {
	gorm.Model
	ExternalID string `json:"external_id"`
	UserID     uint   `json:"user_id"`
	Interface  string `json:"interface"`
}

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Mode      string `json:"mode"`
}

func (u *User) MustGetExternalID(ctx context.Context, interf string) *ExternalID {
	var eid ExternalID

	if tx := db.WithContext(ctx).First(&eid, "user_id = ? AND interface = ?", u.ID, interf); tx.Error == gorm.ErrRecordNotFound {
		eid = ExternalID{
			UserID:    u.ID,
			Interface: interf,
		}
		if err := db.WithContext(ctx).Create(&eid).Error; err != nil {
			panic(err)
		}
	} else if tx.Error != nil {
		panic(tx.Error)
	}

	return &eid
}

func (u *User) AddExternalID(ctx context.Context, externalId, interf string) error {
	var eid ExternalID

	if tx := db.WithContext(ctx).First(&eid, "user_id = ? AND interface = ?", u.ID, interf); tx.Error == gorm.ErrRecordNotFound {
		eid = ExternalID{
			ExternalID: externalId,
			UserID:     u.ID,
			Interface:  interf,
		}
		if err := db.WithContext(ctx).Create(&eid).Error; err != nil {
			return err
		}
	} else if tx.Error != nil {
		return tx.Error
	}

	eid.ExternalID = externalId

	return db.WithContext(ctx).Save(&eid).Error
}

func (u *User) SetMode(ctx context.Context, mode string) error {
	return db.WithContext(ctx).Model(u).Update("mode", mode).Error
}

func (u *User) Save(ctx context.Context) error {
	return db.WithContext(ctx).Save(u).Error
}

func (u *User) LoadMessages(ctx context.Context) ([]*Message, error) {
	var messages []*Message
	if err := db.WithContext(ctx).Where("user_id = ?", u.ID).Find(&messages).Error; err != nil {
		return nil, err
	}

	return messages, nil
}

func (u *User) DeleteMessages(ctx context.Context) error {
	return db.WithContext(ctx).Where("user_id = ?", u.ID).Delete(&Message{}).Error
}
