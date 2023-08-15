package handlers

import (
	"clown-id/internal/models"
	"clown-id/internal/store"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func RegisterAuthHandlers(router *mux.Router, store store.Store, secret string) {
	router.HandleFunc("/login/", handleLogin(store, secret)).Methods("POST")
	router.HandleFunc("/register/", handleRegister(store)).Methods("POST")
}

type HandleLoginRequest struct {
	Email    string `json:"email,omitempty" example:"aboba@gmail.ru"` // either email or username should not be empty
	Username string `json:"username,omitempty" example:"aboba"`
	Password string `json:"password" example:"qwerty123456"`
	AppId    string `json:"app_id" example:"1"`
	ClientId string `json:"client_id" example:"1"`
}

type HandleLoginResponse struct {
	RefreshToken models.RefreshToken `json:"refresh_token"`
}

// Login godoc
// @Summary Аутентификация пользователя по логину или email и паролю
// @Description Необходим либо логин, либо email.
// @Description В случае, если предоставлены оба поля приоритет будет у логина.
// @Description Возвращает пару токенов - access и refresh токен или json с ошибкой
// @Tags Auth
// @ID auth-login
// @Produce json
// @Consume json
// @Param Request body HandleLoginRequest true "json запроса:"
// @Success 200 {object} HandleLoginResponse
// @Failure 400	{object} HttpErrorResponse
// @Router /login/ [post]
func handleLogin(store store.Store, secret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &HandleLoginRequest{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}

		var user *models.User
		var err error

		//TODO: fix beautify this if tree (dive errors to repository)
		if req.Username != "" {
			user, err = store.User().FindByUsername(req.Username)
			if err != nil {
				respondError(w, r, http.StatusBadRequest, err)
				return
			}
		} else if req.Email != "" {
			user, err = store.User().FindByEmail(req.Email)
			if err != nil {
				respondError(w, r, http.StatusBadRequest, err)
				return
			}
		} else {
			respondError(w, r, http.StatusBadRequest, fmt.Errorf("both username and email are empty"))
			return
		}

		if !user.ComparePassword(req.Password) {
			respondError(w, r, http.StatusBadRequest, fmt.Errorf("wrong password"))
			return
		}

		// TODO: check if app id and client id exists (maybe in repository?)
		refreshToken := models.RefreshToken{
			Token:     uuid.New().String(),
			AppId:     req.AppId,
			ClientId:  req.ClientId,
			ExpiresAt: time.Now().AddDate(0, 1, 0).Unix(), // one month
			UserId:    user.ID,
		}
		
		
		if err := store.Token().Create(&refreshToken); err != nil {
			respondError(w, r, http.StatusInternalServerError, err)
		}

		respond(w, r, http.StatusOK, HandleLoginResponse{refreshToken})
	}
}

type HandleRegisterRequest struct {
	Username string `json:"username" example:"aboba"`
	Email    string `json:"email" example:"aboba@gmail.com"`
	Password string `json:"password" example:"aboba32"`
}

// Register godoc
// @Summary Регистрация пользователя.
// @Description Принимает json с пользователем.
// @Description Возвращает либо созданного пользователя либо json с ошибкой.
// @Tags Auth
// @ID auth-register
// @Accept json
// @Produce json
// @Param Request body HandleRegisterRequest true "json пользователя:"
// @Success 200 {object} models.User
// @Failure 400	{object} HttpErrorResponse
// @Router /register/ [post]
func handleRegister(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := &HandleRegisterRequest{}

		if err := json.NewDecoder(r.Body).Decode(request); err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}

		u := &models.User{Username: request.Username, Email: request.Email, Password: request.Password}
		if err := store.User().Create(u); err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}

		u.BeforeSending()
		respond(w, r, http.StatusOK, u)
	}

}
