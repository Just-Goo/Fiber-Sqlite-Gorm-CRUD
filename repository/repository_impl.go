package repository

import "gorm.io/gorm"

type RepositoryImpl struct {
	DB *gorm.DB
}

func (r *RepositoryImpl) Create(model interface{}) interface{} {
	r.DB.Create(model)
	return model
}

// func (r *RepositoryImpl) GetById(model interface{}, id int) []interface{} {

// }

// func (r *RepositoryImpl) GetAll(model interface{}) []interface{} {

// }

// func (r *RepositoryImpl) Update(model interface{}, data interface{}) interface{} {

// }

// func (r *RepositoryImpl) Delete(model interface{}, id int) error {

// }

func NewRepositoryImpl(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{DB: db}
}
