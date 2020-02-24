package controllers

import (
	"chefhub.pw/models"
	"chefhub.pw/views"
)

func NewGalleries(gs models.GalleryService) *Galleries {
	return &Galleries{
		New: views.NewView("bootstrap", "galleries/new"),
		gs:  gs,
	}
}

// Galleries struct holds our views and user service.
type Galleries struct {
	New *views.View
	gs  models.GalleryService
}
