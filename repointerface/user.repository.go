package repointerface

import (
	"go-fiber/model/entity"

	"gorm.io/gorm"
)

type Repository interface {
	AddUser(user entity.User) (entity.User, error)
	GetUsers() ([]entity.User, error)
	GetUser(ID int) (entity.User, error)
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

func (r *repository) GetUsers() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *repository) GetUser(ID int) (entity.User, error) {
	var user entity.User
	err := r.db.Find(&user, ID).Error
	return user, err
}
