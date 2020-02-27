package middleware

import (
	"fmt"
	"net/http"

	"chefhub.pw/context"
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
		cookie, err := r.Cookie("remember_token")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		user, err := mw.UserService.ByRemember(cookie.Value)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		ctx := r.Context()
		ctx = context.WithUser(ctx, user)
		r = r.WithContext(ctx)
		fmt.Println("User found:", user)
		next(w, r)
	})
}