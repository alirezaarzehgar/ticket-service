package model

import (
	"gorm.io/gorm"
)

const (
	USERS_ROLE_SUPER_ADMIN = "super_admin"
	USERS_ROLE_ADMIN       = "admin"
	USERS_ROLE_USER        = "user"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null; unique" json:"user"`
	Email    string `gorm:"not null; unique" json:"email"`
	Password string `gorm:"not null" json:"pass,omitempty"`
	Role     string `gorm:"default:user" json:"role"`
	Blocked  bool   `gorm:"default:false" json:"blocked"`
	Verified bool   `gorm:"default:true" json:"verified"`
}
