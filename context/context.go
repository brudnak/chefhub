package context

import (
	"context"

	"chefhub.pw/models"
)

const (
	userKey privateKey = "user"
)

type privateKey string

// WithUser context
func WithUser(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

// User context
func User(ctx context.Context) *models.User {
	if temp := ctx.Value(userKey); temp != nil {
		if user, ok := temp.(*models.User); ok {
			return user
		}
	}
	return nil
}
