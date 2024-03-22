package main

import (
	"log"

	"github.com/Just-Goo/fibre-sqlite-gorm-crud/controller"
	"github.com/Just-Goo/fibre-sqlite-gorm-crud/database"
)

func main() {
	db, err := database.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}

	controller.InitRouter(db)
}
