package models

import (
	"github.com/deeprave/go-auth/repository"
	"time"
)

type User struct {
	Id              int64     `json:"id"`
	Username        string    `json:"username"`
	Email           string    `json:"email"`
	Given           string    `json:"given"`
	Surname         string    `json:"surname"`
	Phone           string    `json:"phone"`
	IsActive        bool      `json:"is_active"`
	IsAdmin         bool      `json:"is_admin"`
	IsLogindisabled bool      `json:"is_logindisabled"`
	IsVerified      bool      `json:"is_verified"`
	LastLoginAt     time.Time `json:"dt_lastlogin"`
	CreatedAt       time.Time `json:"-"`
	UpdatedAt       time.Time `json:"-"`
}

func (u *User) ScanFields() repository.Ptrs {
	return repository.NewPtrs(
		&u.Id, &u.Username, &u.Email, &u.Given, &u.Surname,
		&u.IsActive, &u.IsAdmin, &u.IsLogindisabled, &u.IsVerified,
		&u.LastLoginAt, &u.CreatedAt, &u.UpdatedAt)
}

var (
	UserTable = repository.NewTable(
		"user",
		"u",
		[]string{
			"id", "username", "email", "given", "surname",
			"is_active", "is_admin", "is_logindisabled", "is_verified",
			"dt_lastlogin", "dt_created", "dt_updated",
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
