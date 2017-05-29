package models

import "time"

// System models

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

type User struct {
	Model
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"-"`
	IsSystem bool   `json:"is_system"`
}
