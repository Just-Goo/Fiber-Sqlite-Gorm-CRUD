package database

import (
	"fmt"

	"github.com/Just-Goo/fibre-sqlite-gorm-crud/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseInstance struct {
	DB *gorm.DB
}

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("store.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	fmt.Println("connected to database successfully")

	db.Logger = logger.Default.LogMode(logger.Info)
	fmt.Println("running migrations")
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	return db, nil
}
