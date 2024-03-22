package handler

import (
	"github.com/Just-Goo/fibre-sqlite-gorm-crud/models"
	"github.com/Just-Goo/fibre-sqlite-gorm-crud/service"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Service service.Service
}

// Users
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

// Products
func (h *Handler) CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	h.Service.CreateProduct(&product)
	return c.Status(fiber.StatusCreated).JSON(product)
}

func (h *Handler) GetProducts(c *fiber.Ctx) error {
	var products []models.Product

	h.Service.FetchProducts(&products)
	return c.Status(fiber.StatusOK).JSON(products)
}

func (h *Handler) GetProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("id must be an integer")
	}

	var product models.Product

	if err = h.Service.FetchProductById(&product, id); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(product)
}

func (h *Handler) UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("id must be aan integer")
	}

	var product, updatedProduct models.Product

	if err := c.BodyParser(&updatedProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	h.Service.UpdateProduct(&product, &updatedProduct, id)
	return c.Status(fiber.StatusOK).JSON(product)
}

func (h *Handler) DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("id must be aan integer")
	}

	var product models.Product

	if err = h.Service.DeleteProduct(&product, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).SendString("product deleted successfully")
}

// Orders
func (h *Handler) CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	result := h.Service.CreateOrder(&order)
	return c.Status(fiber.StatusCreated).JSON(result)
}

func (h *Handler) GetOrders(c *fiber.Ctx) error {
	var orders []models.Order

	result := h.Service.FetchOrders(&orders)
	return c.Status(fiber.StatusOK).JSON(result)
}

func (h *Handler) GetOrderById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("id must be an integer")
	}

	var order models.Order

	if err = h.Service.FetchOrderById(&order, id); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(order)
}

func (h *Handler) UpdateOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("id must be aan integer")
	}

	var order, updatedOrder models.Order

	if err := c.BodyParser(&updatedOrder); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	h.Service.UpdateOrder(&order, &updatedOrder, id)
	return c.Status(fiber.StatusOK).JSON(order)
}

func (h *Handler) DeleteOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("id must be aan integer")
	}

	var order models.Order

	if err = h.Service.DeleteOrder(&order, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).SendString("order deleted successfully")
}

func NewHandler(service service.Service) *Handler {
	return &Handler{Service: service}
}
