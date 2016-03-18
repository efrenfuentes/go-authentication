package models

import (
	"errors"
	"regexp"
	"time"

	"github.com/efrenfuentes/go-authentication/core/crypt"
)

type User struct {
	ID        uint
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`

	Name            string `json:"name"`
	Email           string `json:"email"`
	EncryptPassword string `json:"-"`

	Groups          []Group `gorm:"many2many:user_groups;"   json:"-"`
}

type NewUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) SetPassword(password string) {
	u.EncryptPassword = crypt.Crypt(password, crypt.RandStringBytes(2))
}

func (u *User) SetEmail(email string) error {
	Re := regexp.MustCompile(`^\w[-._\w]*\w@\w[-._\w]*\w\.\w{2,3}$`)
	isValid := Re.MatchString(email)

	if isValid {
		u.Email = email
		return nil
	} else {
		return errors.New("Email format is invalid " + email)
	}
}

func (u User) Authenticate(email, password string) bool {
	salt := u.EncryptPassword[0:2]

	encrypted := crypt.Crypt(password, salt)

	return (email == u.Email) && (encrypted == u.EncryptPassword)
}
