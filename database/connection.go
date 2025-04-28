package database

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectToDb() (*gorm.DB, error) {
	dsn := os.Getenv("GORM_DBSTRING") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	Database = db

	return db, nil
}
