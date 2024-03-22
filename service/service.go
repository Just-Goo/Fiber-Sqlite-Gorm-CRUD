package service

import "github.com/Just-Goo/fibre-sqlite-gorm-crud/repository"

type Service interface {
	CreateUser(model interface{}) interface{}
}

type MainService struct {
	Service Service
}

func NewMainService(repo *repository.Repo) *MainService {
return &MainService{Service: NewServiceImpl(repo.Repo)}
}