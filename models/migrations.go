package models

import "github.com/efrenfuentes/go-authentication/database"


func Migrations() {
	// User
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
	database.DB.AutoMigrate(&Ability{})
	database.DB.Model(&Ability{}).AddUniqueIndex("idx_ability_name", "name")

	ability := Ability{}
	database.DB.Where("name = ?", "all").First(&ability)

	if ability.ID == 0 {
		ability.Name = "All"
		database.DB.Create(&ability)

		database.DB.Model(&ability).Association("Groups").Append(group)
	}
}

