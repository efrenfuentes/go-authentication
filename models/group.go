package models

import (
	"time"
)

type Group struct {
	ID        uint
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`

	Name        string   `json:"name"`
	Description string   `json:"description"`
	Users       []User   `gorm:"many2many:user_groups;"  json:"users"`
	Clients     []Client `gorm:"many2many:client_groups;"  json:"clients"`
}

type NewGroup struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// HasAbility checks if group has a specific ability
func (g Group) HasAbility(abilityName string) bool {
	for _, client := range g.Clients {
		for _, ability := range client.Abilities {
			if (abilityName == ability.Name) || (ability.Name == "all") {
				return true
			}
		}
	}

	return false
}
