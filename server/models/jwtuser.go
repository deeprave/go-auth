package models

import (
	"fmt"
	"strings"
)

type JWTUser struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email,omitempty"`
	Given    string `json:"given,omitempty"`
	Surname  string `json:"surname,omitempty"`
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
