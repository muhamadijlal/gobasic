package migration

import (
	"fmt"

	"github.com/muhamadijlal/gobasic/database"
	"github.com/muhamadijlal/gobasic/models/entity"
)

func RunMigrate() {
	err := database.DB.AutoMigrate(&entity.User{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Success to migrate")
}
