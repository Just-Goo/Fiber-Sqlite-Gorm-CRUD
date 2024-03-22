package repository

import "gorm.io/gorm"

type Repository interface {
	Create(model interface{}) interface{}
	// GetById(model interface{}, id int) []interface{}
	// GetAll(model interface{}) []interface{}
	// Update(model interface{}, data interface{}) interface{}
	// Delete(model interface{}, id int) error
}

type Repo struct {
	Repo Repository
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{Repo: NewRepositoryImpl(db)}
}
