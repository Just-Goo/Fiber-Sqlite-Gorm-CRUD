package repository

import "gorm.io/gorm"

type Repository interface {
	Create(model interface{})
	GetById(model interface{}, id int)
	GetAll(model interface{})
	Update(model interface{}) 
	Delete(model interface{}) error
}

type Repo struct {
	Repo Repository
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{Repo: NewRepositoryImpl(db)}
}
