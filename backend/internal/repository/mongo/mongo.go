package mongo

import (
	"context"

	"qi-rec/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type HistoryRepo struct {
	client   *mongo.Client
	database *mongo.Database
}

func NewHistoryRepo(ctx context.Context, dbURI, dbName string) (*HistoryRepo, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		return nil, err
	}

	database := client.Database(dbName)
	return &HistoryRepo{client: client, database: database}, nil
}

type SongRecommenderHistory struct {
	UserID int            `json:"user_id" bson:"user_id"`
	Tracks []domain.Track `json:"tracks" bson:"tracks"`
}

func (h *HistoryRepo) AddTrack(ctx context.Context, userID int, track *domain.Track) error {
	collection := h.database.Collection("history")
	filter := bson.M{"user_id": userID}
	update := bson.M{"$push": bson.M{"tracks": track}}

	_, err := collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	return err
}

func (h *HistoryRepo) GetHistoryByUserID(ctx context.Context, userID int) ([]*domain.Track, error) {
	collection := h.database.Collection("history")
	filter := bson.M{"user_id": userID}

	var result SongRecommenderHistory
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	var tracks []*domain.Track
	for _, track := range result.Tracks {
		tracks = append(tracks, &track)
	}

	return tracks, nil
}
