package migration

import (
	"fmt"
	"go-fiber/database"
	"go-fiber/model/entity"
)

func RunMigration() {
	// database.DB.AutoMigrate(&entity.User{}, &next-entity)
	err := database.DB.AutoMigrate(&entity.User{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
