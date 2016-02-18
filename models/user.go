package models

import (
	"time"
	"regexp"
	"errors"
	"github.com/efrenfuentes/go-authentication/core/crypt"
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

type User struct {
	ID               uint
	CreatedAt        time.Time
	UpdatedAt        time.Time

	Name             string
	Email            string
	EncryptPassword  string
}

func (u *User)SetPassword(password string) {
	u.EncryptPassword = crypt.Crypt(password, randStringBytes(2))
}

func (u *User)SetEmail(email string) error {
	Re := regexp.MustCompile(`^\w[-._\w]*\w@\w[-._\w]*\w\.\w{2,3}$`)
	isValid := Re.MatchString(email)

	if isValid {
		u.Email = email
		return nil
	} else {
		return errors.New("Email format is invalid " + email)
	}
}

func (u User)Authenticate(email, password string) bool {
	salt := u.EncryptPassword[0:2]

	encrypted := crypt.Crypt(password, salt)

	return (email == u.Email) && (encrypted == u.EncryptPassword)
}
