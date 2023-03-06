package models

import (
	"github.com/deeprave/go-auth/lib"
	"time"
)

type User struct {
	Id              int64     `json:"id"`
	Username        string    `json:"username"`
	Email           string    `json:"email"`
	Given           string    `json:"given,omitempty"`
	Surname         string    `json:"surname,omitempty"`
	Phone           string    `json:"phone,omitempty"`
	IsActive        bool      `json:"is_active"`
	IsAdmin         bool      `json:"is_admin,omitempty"`
	IsLoginDisabled bool      `json:"is_login_disabled,omitempty"`
	IsVerified      bool      `json:"is_verified"`
	IsMFAEnabled    bool      `json:"is_mfa_enabled,omitempty"`
	LastLoginAt     time.Time `json:"dt_lastlogin"`
	CreatedAt       time.Time `json:"-"`
	UpdatedAt       time.Time `json:"-"`
	DeletedAt       time.Time `json:"-"`
}

func (u *User) ScanFields() lib.Ptrs {
	return lib.NewPtrs(
		&u.Id, &u.Username, &u.Email, &u.Given, &u.Surname,
		&u.IsActive, &u.IsAdmin, &u.IsLoginDisabled, &u.IsVerified, &u.IsMFAEnabled,
		&u.LastLoginAt, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt)
}

var (
	UserTable = lib.NewTable(
		"user",
		"u",
		[]string{
			"id", "username", "email", "given", "surname",
			"is_active", "is_admin", "is_login_disabled", "is_verified", "is_mfa_enabled",
			"dt_lastlogin", "dt_created", "dt_updated", "dt_deleted",
		},
	)
)

func (u *User) ToJWTUser() *JWTUser {
	return &JWTUser{
		ID:       u.Id,
		Username: u.Username,
		Email:    u.Email,
		Given:    u.Given,
		Surname:  u.Surname,
	}
}
