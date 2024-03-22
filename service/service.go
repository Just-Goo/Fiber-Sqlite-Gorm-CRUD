package service

import (
	"github.com/Just-Goo/fibre-sqlite-gorm-crud/models"
	"github.com/Just-Goo/fibre-sqlite-gorm-crud/repository"
)

type Service interface {
	CreateUser(model *models.User)
	FetchUsers(model *[]models.User)
	FetchUserById(model *models.User, id int) error
	UpdateUser(model, updatedUser *models.User, id int)
	DeleteUser(model *models.User, id int) error
}

type MainService struct {
	Service Service
}

func NewMainService(repo *repository.Repo) *MainService {
return &MainService{Service: NewServiceImpl(repo.Repo)}
}