package repository

import (
	"context"

	"qi-rec/internal/domain"
	"qi-rec/internal/repository/postgres"
	"qi-rec/internal/repository/postgres/queries"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	CreateUser(ctx context.Context, email string, password string) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	ExistsByID(ctx context.Context, id int) (bool, error)
}

func NewUserRepository(pgxPool *pgxpool.Pool) UserRepository {
	return &postgres.UserRepo{
		Queries: queries.New(pgxPool),
		Pool:    pgxPool,
	}
}

type HistoryRepository interface {
	AddTrack(ctx context.Context, userID int, track *domain.Track) error
	GetHistoryByUserID(ctx context.Context, userID int) ([]*domain.Track, error)
}
