package model

import (
	"gorm.io/gorm"
)

const (
	USERS_ROLE_SUPER_ADMIN = "super_admin"
	USERS_ROLE_ADMIN       = "admin"
	USERS_ROLE_USER        = "user"
)

type OrgAdmin struct {
	OrganizationID, UserID uint
}

type User struct {
	gorm.Model
	Username      string         `gorm:"not null; unique" json:"username"`
	Email         string         `gorm:"not null; unique" json:"email"`
	Password      string         `gorm:"not null" json:"password,omitempty"`
	Role          string         `gorm:"default:user" json:"role"`
	Blocked       bool           `gorm:"default:false" json:"blocked"`
	Verified      bool           `gorm:"default:true" json:"verified"`
	Organizations []Organization `gorm:"many2many:org_admins"`
}

type Organization struct {
	gorm.Model
	Name        string `gorm:"not null; unique" json:"name"`
	Address     string `gorm:"not null" json:"address"`
	PhoneNumber string `gorm:"not null" json:"phone_number"`
	WebsiteUrl  string `json:"website_url"`
	Admins      []User `gorm:"many2many:org_admins"`
	Tickets     []Ticket
}

type Ticket struct {
	UserID         uint   `gorm:"not null" json:"user_id"`
	OrganizationID uint   `gorm:"not null" json:"org_id"`
	Title          string `gorm:"not null" json:"title"`
	Body           string `gorm:"not null" json:"body"`
	AttachmentUrl  string `json:"attachment_url"`
	Seen           bool   `gorm:"default:false" json:"seen"`
	User           *User
	Organization   *Organization
}
