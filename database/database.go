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
	Database.AutoMigrate(&models.User{})
	Database.Model(&models.User{}).AddUniqueIndex("idx_user_email", "email")
}
