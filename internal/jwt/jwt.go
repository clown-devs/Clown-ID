package jwt

import (
	"clown-id/internal/models"
	"clown-id/internal/store"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenPair struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciJ9.eyX0228.LQpEZvladOSc"`
	RefreshToken string `json:"refresh_token" example:"ea7f64d0-9e7a-41ac-a9a3-ca27ee71434f"`
}

// returns access and refresh token
func IssueTokenPair(store store.Store, user *models.User, appId, clientId, secret string) (TokenPair, error) {
	// TODO: check if app id and client id exists (maybe in repository?)
	refreshToken := models.RefreshToken{
		Token:     uuid.New().String(),
		AppId:     appId,
		ClientId:  clientId,
		ExpiresAt: time.Now().AddDate(0, 1, 0).Unix(), // one month
		UserId:    user.ID,
	}

	if err := store.Token().Create(&refreshToken); err != nil {
		return TokenPair{}, err
	}

	payload := jwt.MapClaims{
		"user": user,
		"exp":  time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	accessToken, err := token.SignedString([]byte(secret))

	if err != nil {
		return TokenPair{}, fmt.Errorf("jwt signing error: %e", err)
	}

	return TokenPair{accessToken, refreshToken.Token}, nil
}

func IssueByRefreshToken(refreshToken string, store store.Store, secret string) (TokenPair, error) {
	token, err := store.Token().Find(refreshToken)
	if err != nil {
		return TokenPair{}, err
	}

	user, err := store.User().Find(token.UserId)
	if err != nil {
		return TokenPair{}, err
	}

	return IssueTokenPair(store, user, token.AppId, token.ClientId, secret)
}
