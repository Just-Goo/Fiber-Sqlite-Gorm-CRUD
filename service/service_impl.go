package service

import (
	"fmt"

	"github.com/Just-Goo/fibre-sqlite-gorm-crud/models"
	"github.com/Just-Goo/fibre-sqlite-gorm-crud/repository"
)

type ServiceImpl struct {
	Repo repository.Repository
}

// Users
func (s *ServiceImpl) CreateUser(model *models.User) {
	s.Repo.Create(model)
}

func (s *ServiceImpl) FetchUsers(model *[]models.User) {
	s.Repo.GetAll(model)
}

func (s *ServiceImpl) FetchUserById(model *models.User, id int) error {
	s.Repo.GetById(model, id)
	if model.ID == 0 {
		return fmt.Errorf("user does not exist")
	}
	return nil
}

func (s *ServiceImpl) UpdateUser(model, updatedUser *models.User, id int) {

	// Get the user from database
	s.Repo.GetById(model, id)

	model.FirstName = updatedUser.FirstName
	model.LastName = updatedUser.LastName

	s.Repo.Update(model)
}

func (s *ServiceImpl) DeleteUser(model *models.User, id int) error {

	// Get the user from database
	s.Repo.GetById(model, id)

	// delete the user
	return s.Repo.Delete(model)
}

// Products
func (s *ServiceImpl) CreateProduct(model *models.Product) {
	s.Repo.Create(model)
}

func (s *ServiceImpl) FetchProducts(model *[]models.Product) {
	s.Repo.GetAll(model)
}

func (s *ServiceImpl) FetchProductById(model *models.Product, id int) error {
	s.Repo.GetById(model, id)
	if model.ID == 0 {
		return fmt.Errorf("product does not exist")
	}
	return nil
}

func (s *ServiceImpl) UpdateProduct(model, updatedProduct *models.Product, id int) {

	// Get the product from database
	s.Repo.GetById(model, id)

	model.Name = updatedProduct.Name
	model.SerialNumber = updatedProduct.SerialNumber

	s.Repo.Update(model)
}

func (s *ServiceImpl) DeleteProduct(model *models.Product, id int) error {

	// Get the product from database
	s.Repo.GetById(model, id)

	// delete the product
	return s.Repo.Delete(model)
}

// Orders
func (s *ServiceImpl) CreateOrder(model *models.Order) map[string]interface{} {

	// Get the user
	var user models.User
	s.Repo.GetById(&user, int(model.UserRefer))

	// get the product
	var product models.Product
	s.Repo.GetById(&product, int(model.ProductRefer))

	// create product
	s.Repo.Create(model)

	data := make(map[string]interface{})
	data["id"] = model.ID
	data["user"] = user
	data["product"] = product
	data["order_date"] = model.CreatedAt

	return data
}

func (s *ServiceImpl) FetchOrders(model *[]models.Order) []models.Order {
	s.Repo.GetAll(model)
	var result []models.Order

	// get the user and product associated with an order
	for _, order := range *model {
		var user models.User
		var product models.Product
		var singleOrder models.Order
		s.Repo.GetById(&user, int(order.UserRefer))
		s.Repo.GetById(&product, int(order.ProductRefer))
		singleOrder = models.Order{
			ID: order.ID, ProductRefer: order.ProductRefer,
			Product: product, UserRefer: order.UserRefer,
			User: user, CreatedAt: order.CreatedAt,
		}
		result = append(result, singleOrder)
	}

	return result
}

func (s *ServiceImpl) FetchOrderById(model *models.Order, id int) error {
	s.Repo.GetById(model, id)
	if model.ID == 0 {
		return fmt.Errorf("order does not exist")
	}

	// Get the user
	var user models.User
	s.Repo.GetById(&user, int(model.UserRefer))

	// get the product
	var product models.Product
	s.Repo.GetById(&product, int(model.ProductRefer))

	model.User = user
	model.Product = product

	return nil
}

func (s *ServiceImpl) UpdateOrder(model, updatedOrder *models.Order, id int) {

	// Get the Order from database
	s.Repo.GetById(model, id)

	model.ProductRefer = updatedOrder.ProductRefer
	model.UserRefer = updatedOrder.UserRefer

	// Get the user
	var user models.User
	s.Repo.GetById(&user, int(model.UserRefer))

	// get the product
	var product models.Product
	s.Repo.GetById(&product, int(model.ProductRefer))

	model.User = user
	model.Product = product

	s.Repo.Update(model)
}

func (s *ServiceImpl) DeleteOrder(model *models.Order, id int) error {

	// Get the order from database
	s.Repo.GetById(model, id)

	// delete the order
	return s.Repo.Delete(model)
}

func NewServiceImpl(repo repository.Repository) *ServiceImpl {
	return &ServiceImpl{Repo: repo}
}
