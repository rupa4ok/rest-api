package Models

import (
	"github.com/google/uuid"
)

type Book struct {
	ID       uuid.UUID `gorm:"primary_key;type:char(36)"`
	Name     string    `json:"name"`
	Author   string    `json:"author"`
	Category string    `json:"category"`
}

func (b *Book) TableName() string {
	return "books"
}
