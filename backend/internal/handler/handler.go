package handler

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"qi-rec/internal/handlergen"
	"qi-rec/internal/service/recommendation/spotify"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	spotify *spotify.Client
}

func (s Handler) PostLogout(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s Handler) PostRecommendation(c *gin.Context) { // TODO implement me
	var body handlergen.PostRecommendationJSONRequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"invalid request:": err.Error()})
		return
	}

	uri := *body.PlaylistLink

	id, err := extractIDFromURI(uri)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ids, err := s.spotify.GetTracksByPlaylistID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	t, err := s.spotify.GetTrack(ids[0])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, t)
}

// TODO move to service layer
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

func (s Handler) GetRecommendationHistory(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s Handler) PostSignin(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s Handler) PostSignup(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewHandler(client *spotify.Client) *Handler {
	return &Handler{spotify: client}
}
