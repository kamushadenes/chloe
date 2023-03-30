package memory

import (
	"gorm.io/gorm"
)

type APIKey struct {
	gorm.Model
	User   *User  `json:"user"`
	UserID uint   `json:"userId,omitempty"`
	Hash   string `json:"hash"`
	Key    string `json:"key" gorm:"-"`
}

func NewAPIKey(user *User) *APIKey {
	key := generateSecureToken(32)

	hash := hashString(key)

	return &APIKey{
		User: user,
		Key:  key,
		Hash: hash,
	}
}

func GetAPIKey(key string) (*APIKey, error) {
	var apiKey APIKey

	hash := hashString(key)

	if err := db.Preload("User").Where("hash = ?", hash).First(&apiKey).Error; err != nil {
		return nil, err
	}
	return &apiKey, nil
}
