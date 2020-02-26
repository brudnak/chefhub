package middleware

import (
	"fmt"
	"net/http"
	"time"

	"chefhub.pw/models"
)

// RequireUser struct holds our data
type RequireUser struct {
	models.UserService
}

// Apply helps serve ApplyFn
func (mw *RequireUser) Apply(next http.Handler) http.HandlerFunc {
	return mw.ApplyFn(next.ServeHTTP)
}

// ApplyFn applies our middleware to check if users are logged in
func (mw *RequireUser) ApplyFn(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		fmt.Println("Request started at:", t)

		next(w, r)

		fmt.Println("Request ended at at:", time.Since(t))
	})
}
