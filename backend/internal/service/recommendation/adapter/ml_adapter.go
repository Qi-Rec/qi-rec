package adapter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"qi-rec/internal/domain"
)

type MLAdapter struct {
	host string
	port string
}

func NewAdapter(host, port string) *MLAdapter {
	return &MLAdapter{host: host, port: port}
}

const predictURL = "http://%s:%s/predict"

var ErrUnexpectedStatusFromMLService = fmt.Errorf("unexpected status code from ML service")

func (a *MLAdapter) GetRecommendation(features []domain.TrackFeatures) (string, error) {
	jsonData, err := json.Marshal(struct {
		Features []domain.TrackFeatures `json:"songs"`
	}{
		Features: features,
	})
	if err != nil {
		return "", fmt.Errorf("failed to marshal features: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf(predictURL, a.host, a.port), bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%w: %d", ErrUnexpectedStatusFromMLService, resp.StatusCode)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var id struct {
		ID string `json:"id"`
	}
	if err := json.Unmarshal(body, &id); err != nil {
		return "", fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return id.ID, nil
}
