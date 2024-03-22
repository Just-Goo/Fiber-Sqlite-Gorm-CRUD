package service

import "github.com/Just-Goo/fibre-sqlite-gorm-crud/repository"

type ServiceImpl struct {
	Repo repository.Repository
}

func (s *ServiceImpl) CreateUser(model interface{}) interface{} {
	return s.Repo.Create(model)
}

func NewServiceImpl(repo repository.Repository) *ServiceImpl {
	return &ServiceImpl{Repo: repo}
}
