package handler

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"qi-rec/internal/domain"
	"qi-rec/internal/handlergen"
	"qi-rec/internal/service/recommendation/adapter"
	"qi-rec/internal/service/recommendation/recommend"
	"qi-rec/internal/service/recommendation/spotify"
	"qi-rec/internal/service/user"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const expireDeadline = time.Hour * 24

type Handler struct {
	rec       recommend.Recommender
	us        user.Service
	jwtSecret string
}

func NewHandler(rec recommend.Recommender, us user.Service, jwtSecret string) *Handler {
	return &Handler{rec: rec, us: us, jwtSecret: jwtSecret}
}

func (h *Handler) PostLogout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "", "", true, true)
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func (h *Handler) PostRecommendation(c *gin.Context) {
	var body handlergen.PostRecommendationJSONRequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"invalid request:": err.Error()})
		return
	}

	//var userID int
	//if !setUserID(c, &userID) {
	//	c.JSON(http.StatusUnauthorized, gin.H{"error": "user id not found in token"})
	//	return
	//}

	track, err := h.rec.Recommend(*body.PlaylistLink)
	if err != nil {
		handleError(c, err)
	}

	//err = h.us.AddTrackToHistory(c, userID, track)
	//if err != nil {
	//	handleError(c, err)
	//	return
	//}

	c.JSON(http.StatusOK, track)
}

func (h *Handler) GetRecommendationHistory(c *gin.Context) {
	//var userID int
	//if !setUserID(c, &userID) {
	//	c.JSON(http.StatusUnauthorized, gin.H{"error": "user id not found in token"})
	//	return
	//}
	//
	//history, err := h.us.GetUserHistory(c, userID)
	//if err != nil {
	//	handleError(c, err)
	//	return
	//}

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) PostSignin(c *gin.Context) {
	var body handlergen.PostSigninJSONRequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"invalid request:": err.Error()})
		return
	}

	_, err := h.us.SignIn(c, string(*body.Email), *body.Password)
	if err != nil {
		handleError(c, err)
		return
	}

	//if h.setTokenToCookie(c, err, u) {
	//	c.JSON(http.StatusUnauthorized, gin.H{"error": "failed to set token"})
	//	return
	//}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully signed in"})
}

func (h *Handler) PostSignup(c *gin.Context) {
	var body handlergen.PostSignupJSONRequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"invalid request:": err.Error()})
		return
	}

	_, err := h.us.SignUp(c, string(*body.Email), *body.Password)
	if err != nil {
		handleError(c, err)
		return
	}

	//if h.setTokenToCookie(c, err, u) {
	//	c.JSON(http.StatusUnauthorized, gin.H{"error": "failed to set token"})
	//	return
	//}

	c.JSON(http.StatusCreated, gin.H{"message": "Successfully signed up"})
}

func (h *Handler) setTokenToCookie(c *gin.Context, err error, u *domain.User) bool {
	token, err := h.generateJWT(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return true
	}

	setCookieToken(c, token)
	return false
}

func (h *Handler) generateJWT(u *domain.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": u.ID,
		"exp": time.Now().UTC().Add(expireDeadline).Unix(),
	})

	return token.SignedString([]byte(h.jwtSecret))
}

func setCookieToken(c *gin.Context, token string) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*24, "", "", true, false)
}

func handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, user.ErrUserExists):
		c.JSON(http.StatusConflict, gin.H{"already exists": err.Error()})
	case errors.Is(err, user.ErrInvalidEmail):
		c.JSON(http.StatusBadRequest, gin.H{"invalid email": err.Error()})
	case errors.Is(err, user.ErrEmptyPassword):
		c.JSON(http.StatusBadRequest, gin.H{"empty password": err.Error()})
	case errors.Is(err, user.ErrUserNotFound):
		c.JSON(http.StatusNotFound, gin.H{"not found": err.Error()})
	case errors.Is(err, user.ErrorInvalidPassword):
		c.JSON(http.StatusUnauthorized, gin.H{"invalid password": err.Error()})
	case errors.Is(err, spotify.ErrSpotifyNotFound):
		c.JSON(http.StatusNotFound, gin.H{"not found": err.Error()})
	case errors.Is(err, spotify.ErrEmptyToken):
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	case errors.Is(err, adapter.ErrUnexpectedStatusFromMLService):
		c.JSON(http.StatusInternalServerError, gin.H{"unexpected error from ML service": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func setUserID(c *gin.Context, id *int) bool {
	userIdClaim, ok := c.Get("user_id")
	if !ok {
		return false
	}
	var err error
	*id, err = strconv.Atoi(userIdClaim.(string))
	if err != nil {
		return false
	}

	return true
}
