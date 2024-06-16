package spotify

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"qi-rec/internal/domain"
)

const trackURL = "https://api.spotify.com/v1/tracks/%s"

func (c *Client) GetTrack(id string) (*domain.Track, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(trackURL, id), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	setBearerToken(req, c.token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get track: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, ErrSpotifyNotFound
		}
		return nil, fmt.Errorf("failed to get track, Spotify returned: %s", resp.Status)
	}

	return decodeTrack(resp.Body)
}

func decodeTrack(r io.Reader) (*domain.Track, error) {
	var fullTrack struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Album struct {
			Artists []domain.Artist `json:"artists"`
			Images  []struct {
				URL string `json:"url"`
			} `json:"images"`
		} `json:"album"`
	}

	if err := json.NewDecoder(r).Decode(&fullTrack); err != nil {
		return nil, fmt.Errorf("failed to decode track: %w", err)
	}

	track := &domain.Track{
		ID:       fullTrack.ID,
		Name:     fullTrack.Name,
		Artists:  fullTrack.Album.Artists,
		CoverURL: fullTrack.Album.Images[0].URL,
		Link:     fullTrack.ExternalUrls.Spotify,
	}

	return track, nil
}
