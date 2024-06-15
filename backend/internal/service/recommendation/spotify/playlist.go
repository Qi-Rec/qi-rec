package spotify

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const playlistURL = "https://api.spotify.com/v1/playlists/%s/tracks"

func (c *Client) GetTracksByPlaylistID(id string) ([]string, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(playlistURL, id), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	setBearerToken(req, c.token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get playlist: %w", err)
	}
	defer resp.Body.Close()

	return decodePlaylist(resp.Body)
}

func decodePlaylist(r io.Reader) ([]string, error) {
	var fullPlaylist struct {
		Items []struct {
			Track struct {
				ID string `json:"id"`
			} `json:"track"`
		} `json:"items"`
	}

	if err := json.NewDecoder(r).Decode(&fullPlaylist); err != nil {
		return nil, fmt.Errorf("failed to decode playlist: %w", err)
	}

	var trackIDs []string
	for _, item := range fullPlaylist.Items {
		trackIDs = append(trackIDs, item.Track.ID)
	}

	return trackIDs, nil
}
