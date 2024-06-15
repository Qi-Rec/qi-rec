package handler

import (
	"errors"
	"net/http"

	"qi-rec/internal/handlergen"
	"qi-rec/internal/service/recommendation/adapter"
	"qi-rec/internal/service/recommendation/recommend"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	rec recommend.Recommender
}

func (s Handler) PostLogout(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s Handler) PostRecommendation(c *gin.Context) {
	var body handlergen.PostRecommendationJSONRequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"invalid request:": err.Error()})
		return
	}

	track, err := s.rec.Recommend(*body.PlaylistLink)
	if err != nil {
		if errors.Is(err, adapter.ErrUnexpectedStatusFromMLService) { // ML service error
			c.JSON(http.StatusInternalServerError, gin.H{"Ml service error": err.Error()})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, track)
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

func NewHandler(rec recommend.Recommender) *Handler {
	return &Handler{rec: rec}
}
