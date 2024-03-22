package controller

import (
	"log"

	"github.com/Just-Goo/fibre-sqlite-gorm-crud/controller/handler"
	"github.com/Just-Goo/fibre-sqlite-gorm-crud/repository"
	"github.com/Just-Goo/fibre-sqlite-gorm-crud/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) {

	app := fiber.New()

	repo := repository.NewRepo(db)
	service := service.NewMainService(repo)
	handler.RegisterRoutes(app, service)

	app.Get("/welcome", func(c *fiber.Ctx) error {
		return c.SendString("Welcome")
	})

	log.Fatal(app.Listen(":8080"))
}
