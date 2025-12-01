package entities

import "gorm.io/gorm"

type TodoItem struct {
	gorm.Model

	text string
}
