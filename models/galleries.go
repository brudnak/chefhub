package models

import "github.com/jinzhu/gorm"

// Gallery is our image container resource that visitors
// view.
type Gallery struct {
	gorm.Model
	UserID uint   `gorm:"noy_null;index"`
	Title  string `gorm:"not_null"`
}

// GalleryService is our interface.
type GalleryService interface{}

// GalleryDB interface for our data connection
type GalleryDB interface {
	Create(gallery *Gallery) error
}
