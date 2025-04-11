package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Todo represents a task to be completed
type Todo struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed" gorm:"default:false"`
	DueDate     time.Time `json:"due_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TableName defines the table name for the Todo model
func (Todo) TableName() string {
	return "todos"
}
