package repository

import (
	"context"

	"qi-rec/internal/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, email string, password string) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
}
