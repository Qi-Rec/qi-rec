package spotify

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const tokenURL = "https://accounts.spotify.com/api/token"

var (
	ErrEmptyToken      = fmt.Errorf("empty access token")
	ErrSpotifyNotFound = fmt.Errorf("entity not found")
)

type Client struct {
	token string
}

func NewClient(clientID, clientSecret string) (*Client, error) {
	token, err := getAccessToken(clientID, clientSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	return &Client{token: token}, nil
}

type tokenResponse struct {
	AccessToken string `json:"access_token"`
}

func getAccessToken(clientID, clientSecret string) (string, error) {
	resp, err := http.PostForm(tokenURL, tokenRequestDataValues(clientID, clientSecret))
	if err != nil {
		return "", fmt.Errorf("failed to get access token: %w", err)
	}
	defer resp.Body.Close()

	var token tokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return "", fmt.Errorf("failed to decode token: %w", err)
	}

	if len(token.AccessToken) == 0 {
		return "", ErrEmptyToken
	}

	return token.AccessToken, nil
}

func setBearerToken(req *http.Request, token string) {
	req.Header.Set("Authorization", "Bearer "+token)
}

func tokenRequestDataValues(clientID, clientSecret string) url.Values {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)

	return data
}
