package mongodb

import (
	"context"

	"qi-rec/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type HistoryRepo struct {
	Client   *mongo.Client
	Database *mongo.Database
}

type SongRecommenderHistory struct {
	UserID int            `json:"user_id" bson:"user_id"`
	Tracks []domain.Track `json:"tracks" bson:"tracks"`
}

func (h *HistoryRepo) AddTrack(ctx context.Context, userID int, track *domain.Track) error {
	collection := h.Database.Collection("history")
	filter := bson.M{"user_id": userID}
	update := bson.M{"$push": bson.M{"tracks": track}}

	_, err := collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	return err
}

func (h *HistoryRepo) GetHistoryByUserID(ctx context.Context, userID int) ([]*domain.Track, error) {
	collection := h.Database.Collection("history")
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
