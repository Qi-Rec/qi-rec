package mongo

import (
	"context"

	"qi-rec/internal/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type HistoryRepo struct {
	client   *mongo.Client
	database *mongo.Database
}

func NewHistoryRepo(ctx context.Context, dbURI, dbName string) (*HistoryRepo, error) {
	// TODO implement me
	panic("implement me")
}

type SongRecommenderHistory struct {
	UserID int            `json:"user_id" bson:"user_id"`
	Tracks []domain.Track `json:"tracks" bson:"tracks"`
}

func (h *HistoryRepo) AddTrack(ctx context.Context, userID int, track *domain.Track) error {
	// TODO implement me
	panic("implement me")
}

func (h *HistoryRepo) GetHistoryByUserID(ctx context.Context, userID int) ([]*domain.Track, error) {
	// TODO implement me
	panic("implement me")
}
