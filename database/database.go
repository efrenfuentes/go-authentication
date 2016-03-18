package database

import (
	"github.com/efrenfuentes/go-authentication/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Database gorm.DB

func Init(settings map[string]interface{}) error {
	driver := settings["driver"].(string)
	connection := settings["user"].(string) + ":" + settings["password"].(string) + "@/" + settings["database"].(string)

	db, err := gorm.Open(driver, connection+"?charset=utf8&parseTime=True")

	if err != nil {
		return err
	}

	Database = db

	migrations()

	return nil
}

func migrations() {
	// User
	Database.AutoMigrate(&models.User{})
	Database.Model(&models.User{}).AddUniqueIndex("idx_user_email", "email")

	user := models.User{}
	Database.Where("email = ?", "admin@authenticate.com").First(&user)

	if user.ID == 0 {
		user.Name = "Administrator"
		user.SetEmail("admin@authenticate.com")
		user.SetPassword("admin")
		Database.Create(&user)
	}

	// Group
	Database.AutoMigrate(&models.Group{})
	Database.Model(&models.Group{}).AddUniqueIndex("idx_group_name", "name")

	group := models.Group{}
	Database.Where("name = ?", "Administrators").First(&group)

	if group.ID == 0 {
		group.Name = "Administrators"
		group.Description = "Administrators of authentication"
		Database.Create(&group)

		Database.Model(&group).Association("Users").Append(user)
	}

	// Client
	Database.AutoMigrate(&models.Client{})
	Database.Model(&models.Client{}).AddUniqueIndex("idx_client_name", "name")

	client := models.Client{}
	Database.Where("name = ?", "Go-Authenticate").First(&client)

	if client.ID == 0 {
		client.Name = "Go-Authenticate"
		client.GenerateKeys()
		Database.Create(&client)

		Database.Model(&client).Association("Groups").Append(group)
	}

	// Ability
	Database.AutoMigrate(&models.Ability{})
	Database.Model(&models.Ability{}).AddUniqueIndex("idx_ability_name", "name")

	ability := models.Ability{}
	Database.Where("name = ?", "all").First(&ability)

	if ability.ID == 0 {
		ability.Name = "All"
		Database.Create(&ability)

		Database.Model(&ability).Association("Groups").Append(group)
	}
}
