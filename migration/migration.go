package migration

import (
	"fmt"
	"gin-scaffold/model"
	"gin-scaffold/pkg/database"
)

func Migration() {
	err := database.DB.AutoMigrate(
		&model.Account{},
	)
	if err != nil {
		fmt.Println("database migration is failed!")
	}
}
