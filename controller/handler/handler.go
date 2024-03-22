package handler

import (
	"github.com/Just-Goo/fibre-sqlite-gorm-crud/models"
	"github.com/Just-Goo/fibre-sqlite-gorm-crud/service"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Service service.Service
}

func (h *Handler) CreateUser(c *fiber.Ctx) error  {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	createdUser := h.Service.CreateUser(&user)
	return c.Status(201).JSON(createdUser)
}

func NewHandler(service service.Service) *Handler {
	return &Handler{Service: service}
}
