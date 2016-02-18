package database


import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/efrenfuentes/go-authentication/models"
)

var database gorm.DB

func Init(settings map[string]interface{}) error {
	driver := settings["driver"].(string)
	connection := settings["user"].(string) + ":" + settings["password"].(string) + "@/" + settings["database"].(string)

	db, err := gorm.Open(driver, connection)

	if err != nil {
		return err
	}

	database = db

	migrations()

	return nil
}

func migrations() {
	database.AutoMigrate(&models.User{})
}