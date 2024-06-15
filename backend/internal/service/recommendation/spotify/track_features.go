package spotify

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"qi-rec/internal/domain"
)

const audioFeaturesURL = "https://api.spotify.com/v1/audio-features/%s"

func (c *Client) GetTrackFeatures(id string) (*domain.TrackFeatures, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(audioFeaturesURL, id), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	setBearerToken(req, c.token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get track features: %w", err)
	}
	defer resp.Body.Close()

	return decodeTrackFeatures(resp.Body)
}

func decodeTrackFeatures(r io.Reader) (*domain.TrackFeatures, error) {
	var features *domain.TrackFeatures
	if err := json.NewDecoder(r).Decode(&features); err != nil {
		return nil, fmt.Errorf("failed to decode track features: %w", err)
	}

	return features, nil
}
