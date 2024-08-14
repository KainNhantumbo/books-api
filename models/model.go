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
	Name          string    `gorm:"not null" json:"name"`
	Description   string    `gorm:"not null" json:"description"`
	IsAvailable   bool      `gorm:"not null" json:"isAvailable"`
	Publisher     string    `gorm:"not null" json:"publisher"`
	Country       string    `gorm:"not null" json:"country"`
	Category      string    `gorm:"not null" json:"category"`
	PublishedDate time.Time `json:"publishedDate"`
	PageCount     int32     `gorm:"not null" json:"pageCount"`
}

type Image struct {
	Base
	ImageId  string `json:"imageId"`
	ImageUrl string `json:"imageUrl"`
}

type User struct {
	Base
	FirstName string `gorm:"not null" json:"firstName"`
	LastName  string `gorm:"not null" json:"lastName"`
	Email     string `gorm:"uniqueIndex;not null" json:"email"`
	Password  string `gorm:"not null" json:"password"`
	Bio       string `json:"bio"`
	Avatar    *Image `json:"avatar"`
}

type SocialMedia struct {
	Base
	Facebook string `json:"facebook"`
	Blog     string `json:"blog"`
	Github   string `json:"github"`
	LinkedIn string `json:"linkedIn"`
	UserId   *User  `gorm:"not null" json:"userId"`
}
