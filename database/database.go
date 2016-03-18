package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB gorm.DB

func Init(settings map[string]interface{}) error {
	driver := settings["driver"].(string)
	connection := settings["user"].(string) + ":" + settings["password"].(string) + "@/" + settings["database"].(string)

	db, err := gorm.Open(driver, connection+"?charset=utf8&parseTime=True")

	if err != nil {
		return err
	}

	DB = db

	return nil
}