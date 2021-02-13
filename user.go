package model

import (
	"errors"
	"regexp"
	"strings"
)

// errors
var (
	ErrPassTooShort = errors.New("user password is too short")
	ErrInvalidEmail = errors.New("email seems to be in an invalid format")
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// User model
type User struct {
	Email string `json:"email,omitempty"`
	Pass  string `json:"-"`
}

// Validate returns an error if the user is not valid
func (u *User) Validate() error {
	u.Email = strings.ToLower(u.Email)
	if len(u.Pass) < 8 {
		return ErrPassTooShort
	}
	if !emailRegex.MatchString(u.Email) {
		return ErrInvalidEmail
	}
	return nil
}
