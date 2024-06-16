package queries

import (
	"context"
	"fmt"

	"qi-rec/internal/domain"
)

const insertUser = `
INSERT INTO users (email, password)
VALUES ($1, $2)
RETURNING id
`

func (q *Queries) CreateUser(ctx context.Context, email string, password string) (*domain.User, error) {
	user := &domain.User{
		Email:    email,
		Password: password,
	}
	row := q.pool.QueryRow(ctx, insertUser, email, password)
	if err := row.Scan(&user.ID); err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}

	return user, nil
}

const selectUserByEmail = `
SELECT id, email, password
FROM users
WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	row := q.pool.QueryRow(ctx, selectUserByEmail, email)
	if err := row.Scan(&user.ID, &user.Email, &user.Password); err != nil {
		return nil, fmt.Errorf("failed to select user: %w", err)
	}

	return &user, nil
}

const selectUserByID = `
SELECT id, email, password
FROM users
WHERE id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User
	row := q.pool.QueryRow(ctx, selectUserByID, id)
	if err := row.Scan(&user.ID, &user.Email, &user.Password); err != nil {
		return nil, fmt.Errorf("failed to select user: %w", err)
	}

	return &user, nil
}

const existsByEmail = `
SELECT EXISTS (
	SELECT 1
	FROM users
	WHERE email = $1
)
`

func (q *Queries) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var exists bool
	row := q.pool.QueryRow(ctx, existsByEmail, email)
	if err := row.Scan(&exists); err != nil {
		return false, fmt.Errorf("failed to check if user exists: %w", err)
	}

	return exists, nil
}

const existsByID = `
SELECT EXISTS (
	SELECT 1
	FROM users
	WHERE id = $1
)
`

func (q *Queries) ExistsByID(ctx context.Context, id string) (bool, error) {
	var exists bool
	row := q.pool.QueryRow(ctx, existsByID, id)
	if err := row.Scan(&exists); err != nil {
		return false, fmt.Errorf("failed to check if user exists: %w", err)
	}

	return exists, nil
}
