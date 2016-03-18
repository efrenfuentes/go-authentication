package models

import (
	"time"

	"github.com/efrenfuentes/go-authentication/core/crypt"
)

type Client struct {
	ID        uint
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`

	Name            string `json:"name"`
	ClientID        string `json:"client_id"`
	ClientSecret    string `json:"client_secret"`
	Groups          []Group `gorm:"many2many:client_groups;"  json:"-"`
}

type NewClient struct {
	Name            string `json:"name"`
}

func (c *Client) GenerateKeys() {
	c.ClientID = crypt.RandStringBytes(10)
	c.ClientSecret = crypt.RandStringBytes(30)
}