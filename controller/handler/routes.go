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
	v1.Get("/users", handler.GetUsers)
	v1.Get("/users/:id", handler.GetUserById)
	v1.Put("/users/:id", handler.UpdateUser)
	v1.Delete("/users/:id", handler.DeleteUser)

	// Product endpoints
	v1.Post("/products", handler.CreateProduct)
	v1.Get("/products", handler.GetProducts)
	v1.Get("/products/:id", handler.GetProductById)
	v1.Put("/products/:id", handler.UpdateProduct)
	v1.Delete("/products/:id", handler.DeleteProduct)

	// Order endpoints
	v1.Post("/orders", handler.CreateOrder)
	v1.Get("/orders", handler.GetOrders)
	v1.Get("/orders/:id", handler.GetOrderById)
	v1.Put("/orders/:id", handler.UpdateOrder)
	v1.Delete("/orders/:id", handler.DeleteOrder)

}
