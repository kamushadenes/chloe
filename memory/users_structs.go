package memory

import (
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
