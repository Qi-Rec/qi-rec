package repository

import (
	"context"
	"errors"
	"fmt"

	"qi-rec/internal/domain"
	"qi-rec/internal/repository/postgres/queries"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	CreateUser(ctx context.Context, email string, password string) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	ExistsByID(ctx context.Context, id string) (bool, error)
}

type HistoryRepository interface {
}

type userRepo struct {
	*queries.Queries // SQL queries
	pool             *pgxpool.Pool
}

func NewUserRepository(pgxPool *pgxpool.Pool) UserRepository {
	return &userRepo{
		Queries: queries.New(pgxPool),
		pool:    pgxPool,
	}
}

func SetupPgxPool(ctx context.Context, DbURL string) (*pgxpool.Pool, error) {
	pgxConfig, err := pgxpool.ParseConfig(DbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse pgx config: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create new pgx pool with config: %w", err)
	}

	return pool, nil
}

func ProcessMigration(migrationURL string, dbSource string) error {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		return fmt.Errorf("failed to create new migration: %w", err)
	}

	if err = migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to migrate up: %w", err)
	}
	defer migration.Close()

	return nil
}
