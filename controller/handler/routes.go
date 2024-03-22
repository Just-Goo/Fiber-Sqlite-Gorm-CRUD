package handler

import (
	"github.com/Just-Goo/fibre-sqlite-gorm-crud/service"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, service *service.MainService) {
	handler := NewHandler(service.Service)

	v1 := app.Group("/api/v1")

	// User endpoints
	v1.Post("/users", handler.CreateUser)

}
