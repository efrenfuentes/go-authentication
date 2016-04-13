package models

import (
	"errors"
	"regexp"
	"time"

	"github.com/efrenfuentes/go-authentication/core/crypt"
)

// User keep user information like email and password for authentication
type User struct {
	ID        uint
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`

	Name            string `json:"name"`
	Email           string `json:"email"`
	EncryptPassword string `json:"-"`

	Groups []Group `gorm:"many2many:user_groups;"   json:"groups"`
}

// NewUser struct has the information to create a new user
type NewUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthUser struct has the information to be use when an user try to authenticate
type AuthUser struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	ClientSecret string `json:"client_secret"`
}

// SetPassword sets the user password after encrypt it
func (u *User) SetPassword(password string) {
	u.EncryptPassword = crypt.Crypt(password, crypt.RandStringBytes(2))
}

// SetEmail sets the user email after validate it
func (u *User) SetEmail(email string) error {
	Re := regexp.MustCompile(`^\w[-._\w]*\w@\w[-._\w]*\w\.\w{2,3}$`)
	isValid := Re.MatchString(email)

	if isValid {
		u.Email = email
		return nil
	}

	return errors.New("Email format is invalid " + email)
}

// Authenticate check if email and password are correct and if have
// authorization for use the client
func (u User) Authenticate(email, password, secretKey string) bool {
	salt := u.EncryptPassword[0:2]

	encrypted := crypt.Crypt(password, salt)

	return (email == u.Email) && (encrypted == u.EncryptPassword) && u.HasClientAccess(secretKey)
}

// HasClientAccess checks if user has authorization to authenticate in a client
func (u User) HasClientAccess(secretKey string) bool {
	for _, group := range u.Groups {
		for _, client := range group.Clients {
			if secretKey == client.ClientSecret {
				return true
			}
		}
	}

	return false
}

// HasAbility checks if user has a specific ability
func (u User) HasAbility(abilityName string) bool {
	for _, group := range u.Groups {
		if group.HasAbility(abilityName) {
			return true
		}
	}

	return false
}
