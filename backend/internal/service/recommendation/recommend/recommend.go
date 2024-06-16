package recommend

import (
	"fmt"
	"net/url"
	"strings"

	"qi-rec/internal/domain"
	"qi-rec/internal/service/recommendation/adapter"
	"qi-rec/internal/service/recommendation/spotify"
)

type Recommender interface {
	Recommend(playlistURI string) (*domain.Track, error)
}

type recommender struct {
	spotify   *spotify.Client
	mlAdapter *adapter.MLAdapter
}

func (r *recommender) Recommend(playlistURI string) (*domain.Track, error) {
	id, err := extractIDFromURI(playlistURI)
	if err != nil {
		return nil, fmt.Errorf("failed to extract playlist ID: %w", err)
	}

	ids, err := r.spotify.GetTracksByPlaylistID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get tracks by playlist ID: %w", err)
	}

	features := make([]domain.TrackFeatures, len(ids))
	for i := range ids {
		f, err := r.spotify.GetTrackFeatures(ids[i])
		if err != nil {
			return nil, fmt.Errorf("failed to get track features: %w", err)
		}

		features[i] = *f
	}

	recommendedID, err := r.mlAdapter.GetRecommendation(features)
	if err != nil {
		return nil, fmt.Errorf("failed to get recommendation: %w", err)
	}

	t, err := r.spotify.GetTrack(recommendedID)
	if err != nil {
		return nil, fmt.Errorf("failed to get track: %w", err)
	}

	return t, nil
}

func NewRecommender(spotify *spotify.Client, mlAdapter *adapter.MLAdapter) Recommender {
	return &recommender{
		spotify:   spotify,
		mlAdapter: mlAdapter,
	}
}

func extractIDFromURI(uri string) (string, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return "", err
	}

	pathParts := strings.Split(u.Path, "/")
	if len(pathParts) < 3 || pathParts[1] != "playlist" {
		return "", fmt.Errorf("invalid Spotify playlist URL")
	}

	return pathParts[2], nil
}
