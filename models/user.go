package models

import (
	"time"
	"regexp"
	"errors"
)

type User struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time

	Name             string
	email            string
	encrypt_password string
}

func (u *User)SetPassword(password string) {
	u.encrypt_password = password
}

func (u *User)SetEmail(email string) error {
	Re := regexp.MustCompile(`^\w[-._\w]*\w@\w[-._\w]*\w\.\w{2,3}$`)
	isValid := Re.MatchString(email)

	if isValid {
		u.email = email
		return nil
	} else {
		return errors.New("Email format is invalid " + email)
	}
}

func (u User)GetEmail() string {
	return u.email
}

func (u User)Authenticate(email, password string) bool {
	return false
}