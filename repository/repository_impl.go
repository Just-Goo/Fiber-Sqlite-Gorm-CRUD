package repository

import "gorm.io/gorm"

type RepositoryImpl struct {
	DB *gorm.DB
}

func (r *RepositoryImpl) Create(model interface{}) {
	r.DB.Create(model)
}

func (r *RepositoryImpl) GetById(model interface{}, id int) {
	r.DB.Find(model, "id = ?", id)
}

func (r *RepositoryImpl) GetAll(model interface{}) {
	r.DB.Find(model)
}

func (r *RepositoryImpl) Update(model interface{}) {
	r.DB.Save(model)
}

func (r *RepositoryImpl) Delete(model interface{}) error {
	return r.DB.Delete(model).Error
}

func NewRepositoryImpl(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{DB: db}
}
