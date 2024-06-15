package queries

import (
	"context"

	"qi-rec/internal/domain"
)

const insertUser = `
INSERT INTO users (email, password)
VALUES ($1, $2)
`

func (q *Queries) CreateUser(ctx context.Context, email string, password string) (*domain.User, error) {
	panic("")
}
func (q *Queries) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	panic("")
}
func (q *Queries) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	panic("")
}
