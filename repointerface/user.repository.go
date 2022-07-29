package repointerface

import (
	"go-fiber/model/entity"

	"gorm.io/gorm"
)

type Repository interface {
	AddUser(user entity.User) (entity.User, error)
}

type repository struct {
	db *gorm.DB
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db} // agar bisa diakses di main
}

func (r *repository) AddUser(user entity.User) (entity.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}
