package database

import (
	"fmt"
	"go-fiber/models"
	"go-fiber/pkg/mysql"
)

func RunMigration() {
	// database.DB.AutoMigrate(&entity.User{}, &next-entity)
	err := mysql.DB.AutoMigrate(&models.User{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
