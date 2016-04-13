package models

import (
	"fmt"

	"github.com/efrenfuentes/go-authentication/database"
)

// Migrations creates data tables and records for admin user,
// base client, groups and abilities
func Migrations() {
	fmt.Println("Starting migrations...")
	// User
	fmt.Println("    creating table users...")
	database.DB.AutoMigrate(&User{})
	database.DB.Model(&User{}).AddUniqueIndex("idx_user_email", "email")

	user := User{}
	database.DB.Where("email = ?", "admin@authenticate.com").First(&user)

	if user.ID == 0 {
		user.Name = "Administrator"
		user.SetEmail("admin@authenticate.com")
		user.SetPassword("admin")
		database.DB.Create(&user)
	}

	// Group
	fmt.Println("    creating table groups...")
	database.DB.AutoMigrate(&Group{})
	database.DB.Model(&Group{}).AddUniqueIndex("idx_group_name", "name")

	group := Group{}
	database.DB.Where("name = ?", "Administrators").First(&group)

	if group.ID == 0 {
		group.Name = "Administrators"
		group.Description = "Administrators of authentication"
		database.DB.Create(&group)

		database.DB.Model(&group).Association("Users").Append(user)
	}

	// Client
	fmt.Println("    creating table clients...")
	database.DB.AutoMigrate(&Client{})
	database.DB.Model(&Client{}).AddUniqueIndex("idx_client_name", "name")

	client := Client{}
	database.DB.Where("name = ?", "Go-Authenticate").First(&client)

	if client.ID == 0 {
		client.Name = "Go-Authenticate"
		client.GenerateKeys()
		database.DB.Create(&client)

		database.DB.Model(&client).Association("Groups").Append(group)
	}

	// Ability
	fmt.Println("    creating table abilities...")
	database.DB.AutoMigrate(&Ability{})
	database.DB.Model(&Ability{}).AddUniqueIndex("idx_ability_name", "name")

	ability := Ability{}
	database.DB.Where("name = ?", "all").First(&ability)

	if ability.ID == 0 {
		ability.Name = "all"
		database.DB.Create(&ability)

		database.DB.Model(&ability).Association("Clients").Append(client)
	}

	ability = Ability{}
	database.DB.Where("name = ?", "read_users").First(&ability)

	if ability.ID == 0 {
		ability.Name = "read_users"
		database.DB.Create(&ability)
	}

	ability = Ability{}
	database.DB.Where("name = ?", "create_user").First(&ability)

	if ability.ID == 0 {
		ability.Name = "create_user"
		database.DB.Create(&ability)
	}

	ability = Ability{}
	database.DB.Where("name = ?", "update_user").First(&ability)

	if ability.ID == 0 {
		ability.Name = "update_user"
		database.DB.Create(&ability)
	}

	ability = Ability{}
	database.DB.Where("name = ?", "delete_user").First(&ability)

	if ability.ID == 0 {
		ability.Name = "delete_user"
		database.DB.Create(&ability)
	}

	ability = Ability{}
	database.DB.Where("name = ?", "read_clients").First(&ability)

	if ability.ID == 0 {
		ability.Name = "read_clients"
		database.DB.Create(&ability)
	}

	ability = Ability{}
	database.DB.Where("name = ?", "create_client").First(&ability)

	if ability.ID == 0 {
		ability.Name = "create_client"
		database.DB.Create(&ability)
	}

	ability = Ability{}
	database.DB.Where("name = ?", "update_client").First(&ability)

	if ability.ID == 0 {
		ability.Name = "update_client"
		database.DB.Create(&ability)
	}

	ability = Ability{}
	database.DB.Where("name = ?", "delete_client").First(&ability)

	if ability.ID == 0 {
		ability.Name = "delete_client"
		database.DB.Create(&ability)
	}

	ability = Ability{}
	database.DB.Where("name = ?", "read_groups").First(&ability)

	if ability.ID == 0 {
		ability.Name = "read_groups"
		database.DB.Create(&ability)
	}

	ability = Ability{}
	database.DB.Where("name = ?", "create_group").First(&ability)

	if ability.ID == 0 {
		ability.Name = "create_group"
		database.DB.Create(&ability)
	}

	ability = Ability{}
	database.DB.Where("name = ?", "update_group").First(&ability)

	if ability.ID == 0 {
		ability.Name = "update_group"
		database.DB.Create(&ability)
	}

	ability = Ability{}
	database.DB.Where("name = ?", "delete_group").First(&ability)

	if ability.ID == 0 {
		ability.Name = "delete_group"
		database.DB.Create(&ability)
	}

	ability = Ability{}
	database.DB.Where("name = ?", "read_abilities").First(&ability)

	if ability.ID == 0 {
		ability.Name = "read_abilities"
		database.DB.Create(&ability)
	}

	ability = Ability{}
	database.DB.Where("name = ?", "create_ability").First(&ability)

	if ability.ID == 0 {
		ability.Name = "create_ability"
		database.DB.Create(&ability)
	}

	ability = Ability{}
	database.DB.Where("name = ?", "update_ability").First(&ability)

	if ability.ID == 0 {
		ability.Name = "update_ability"
		database.DB.Create(&ability)
	}

	ability = Ability{}
	database.DB.Where("name = ?", "delete_ability").First(&ability)

	if ability.ID == 0 {
		ability.Name = "delete_ability"
		database.DB.Create(&ability)
	}
}
