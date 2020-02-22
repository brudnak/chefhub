package controllers

import "chefhub.pw/views"

// NewStatic returns address of our static pages.
func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bootstrap", "static/home"),
		Contact: views.NewView("bootstrap", "static/contact"),
	}
}

// Static is our data structure to hold static views.
type Static struct {
	Home    *views.View
	Contact *views.View
}
