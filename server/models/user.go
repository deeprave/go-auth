package models

import (
	"errors"
	"fmt"
	"github.com/deeprave/go-auth/repository"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type User struct {
	Id              int64     `json:"id"`
	Username        string    `json:"username"`
	Email           string    `json:"email"`
	Given           string    `json:"given"`
	Surname         string    `json:"surname"`
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

type CredType string

const (
	Plain CredType = "plain"
	Hash  CredType = "hash"
	TOTP  CredType = "totp"
)

func IsPassword(cType CredType) bool {
	return cType == Plain || cType == Hash
}

type Credential struct {
	UserId int64               `json:"user_id"`
	Type   CredType            `json:"type"`
	Data   map[CredType]string `json:"data"`
}

func (c *Credential) ScanFields() repository.Ptrs {
	return repository.NewPtrs(&c.UserId, &c.Type, &c.Data)
}

var (
	CredentialTable = repository.NewTable(
		"credential",
		"c",
		[]string{
			"user_id", "type", "data",
		},
	)
)

type JWTUser struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email,omitempty"`
	Given    string `json:"given,omitempty"`
	Surname  string `json:"surname,omitempty"`
}

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	TOTP     string `json:"totp,omitempty"`
}

type CredentialManager struct {
	// Name of the issuing Organization/Company.
	Issuer string
	// Name of the User's Account (eg, email address)
	AccountName string
}

func NewCredentialManager(issuer string, account string) *CredentialManager {
	return &CredentialManager{
		Issuer:      issuer,
		AccountName: account,
	}
}

func (u *User) NewCredentialManager(issuer string) *CredentialManager {
	return &CredentialManager{
		Issuer:      issuer,
		AccountName: u.Username,
	}
}

func (m *CredentialManager) NewCredential(userId int64, credType CredType, data ...string) (*Credential, error) {
	var (
		stored       = map[CredType]string{}
		err    error = nil
	)

	switch credType {
	case Plain:
		stored[credType] = data[0]
	case Hash:
		var bytevalue []byte
		bytevalue, err = bcrypt.GenerateFromPassword([]byte(data[0]), bcrypt.DefaultCost)
		if err == nil {
			stored[credType] = string(bytevalue)
		}
	case TOTP:
		var key *otp.Key
		key, err = totp.Generate(totp.GenerateOpts{
			Issuer:      m.Issuer,
			AccountName: m.AccountName,
		})
		if err == nil {
			stored[credType] = key.Secret()
		}
	}

	return &Credential{
		UserId: userId,
		Type:   credType,
		Data:   stored,
	}, err
}

func (c *Credential) Matches(phrase string) (bool, error) {
	// extract the expected value, hash or secret
	password, ok := c.Data[c.Type]
	if !ok {
		return false, fmt.Errorf("unexpected error in credential type %v", c.Type)
	}
	switch c.Type {
	// plain password
	case Plain:
		return password == phrase, nil
	// hashed password
	case Hash:
		err := bcrypt.CompareHashAndPassword([]byte(password), []byte(phrase))
		if err != nil {
			if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
				err = nil
			}
			return false, err
		}
		return true, nil
	// one time password
	case TOTP:
		return totp.Validate(phrase, password), nil
	default:
		return false, nil
	}
}

func (r *AuthRequest) Matches(credentials []*Credential) (bool, error) {
	var (
		err      error = nil
		valid          = false
		password       = false
	)
	for _, credential := range credentials {
		secret := r.Password
		if IsPassword(credential.Type) {
			password = true
		} else {
			secret = r.TOTP
		}
		valid, err = credential.Matches(secret)
		if err != nil || !valid {
			break
		}
	}
	// all provided credentials are valid + at least one password type used
	return valid && password, err
}

func (j *JWTUser) LongName() string {
	parts := []string{j.Username}
	if j.Given != "" {
		if j.Surname != "" {
			parts = append(parts, fmt.Sprintf("(%s %s)", j.Given, j.Surname))
		} else {
			parts = append(parts, j.Given)
		}
	} else if j.Surname != "" {
		parts = append(parts, j.Surname)
	}
	if j.Email != "" {
		parts = append(parts, fmt.Sprintf("<%s>", j.Email))
	}
	return strings.Join(parts, " ")
}

func (u *User) ToJWTUser() *JWTUser {
	return &JWTUser{
		ID:       u.Id,
		Username: u.Username,
		Email:    u.Email,
		Given:    u.Given,
		Surname:  u.Surname,
	}
}
