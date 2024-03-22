package service

import (
	"fmt"

	"github.com/Just-Goo/fibre-sqlite-gorm-crud/models"
	"github.com/Just-Goo/fibre-sqlite-gorm-crud/repository"
)

type ServiceImpl struct {
	Repo repository.Repository
}

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

func NewServiceImpl(repo repository.Repository) *ServiceImpl {
	return &ServiceImpl{Repo: repo}
}
