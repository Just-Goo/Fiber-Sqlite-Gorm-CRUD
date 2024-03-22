package handler

import (
	"github.com/Just-Goo/fibre-sqlite-gorm-crud/models"
	"github.com/Just-Goo/fibre-sqlite-gorm-crud/service"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Service service.Service
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	h.Service.CreateUser(&user)
	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *Handler) GetUsers(c *fiber.Ctx) error {
	var users []models.User

	h.Service.FetchUsers(&users)
	return c.Status(fiber.StatusOK).JSON(users)
}

func (h *Handler) GetUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("id must be an integer")
	}

	var user models.User

	if err = h.Service.FetchUserById(&user, id); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("id must be aan integer")
	}

	var user, updatedUser models.User

	if err := c.BodyParser(&updatedUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	h.Service.UpdateUser(&user, &updatedUser, id)
	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("id must be aan integer")
	}

	var user models.User

	if err = h.Service.DeleteUser(&user, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).SendString("user deleted successfully")
}

func NewHandler(service service.Service) *Handler {
	return &Handler{Service: service}
}
