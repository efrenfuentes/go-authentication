package models

import "time"

type Ability struct {
	ID        uint
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`

	Name    string   `json:"name"`
	Clients []Client `gorm:"many2many:ability_clients;"  json:"clients"`
}

type NewAbility struct {
	Name string `json:"name"`
}
