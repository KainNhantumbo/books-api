package model

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt,omitempty"`
}

type Book struct {
	Base
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	IsAvailable   bool      `json:"isAvailable"`
	Publisher     string    `json:"publisher"`
	Country       string    `json:"country"`
	Category      string    `json:"category"`
	PublishedDate time.Time `json:"publishedDate"`
	PageCount     int32     `json:"pageCount"`
}

type User struct {
	Base
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `gorm:"uniqueIndex;not null" json:"email"`
	Password  string `json:"password"`
	Author    string `json:"author"`
}
