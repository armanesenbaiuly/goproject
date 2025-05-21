package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title     string `gorm:"not null" json:"title"`
	Completed bool   `json:"completed"`
}
