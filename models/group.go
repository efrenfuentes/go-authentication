package models

import (
	"time"
)

type Group struct {
	ID        uint
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`

	Name            string `json:"name"`
	Description     string `json:"description"`
	Users           []User `gorm:"many2many:user_groups;"  json:"users"`
	Clients         []Client `gorm:"many2many:client_groups;"  json:"clients"`
	Abilities       []Ability `gorm:"many2many:ability_groups;"  json:"abilities"`
}

type NewGroup struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
}
