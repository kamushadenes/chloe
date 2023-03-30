package memory

import "gorm.io/gorm"

type ApiKey struct {
	gorm.Model
	User   *User  `json:"user"`
	UserID uint   `json:"userId,omitempty"`
	Hash   string `json:"hash"`
}
