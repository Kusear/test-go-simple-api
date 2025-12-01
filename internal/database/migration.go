package database

import (
	"fmt"
	"test-go-simple-api/internal/entities"

	"gorm.io/gorm"
)

func Up(db *gorm.DB) error {
	err := db.AutoMigrate(&entities.User{}, &entities.TodoItem{})
	if err != nil {
		fmt.Printf("Error migrating database: %v\n", err)
		return err
	}

	return nil
}
