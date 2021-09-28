package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DbConnection *gorm.DB

func init() {
	var err error
	DbConnection, err = gorm.Open(sqlite.Open("react_go_authentication.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DbConnection.AutoMigrate(&User{})
}
