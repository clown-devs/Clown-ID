package handlers

import (
	"clown-id/internal/jwt"
	"clown-id/internal/models"
	"clown-id/internal/store"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterAuthHandlers(router *mux.Router, store store.Store, secret string) {
	router.HandleFunc("/register/", handleRegister(store)).Methods("POST")
	router.HandleFunc("/login/", handleLogin(store, secret)).Methods("POST")
	router.HandleFunc("/refresh/", handleRefreshToken(store, secret)).Methods("POST")
	router.HandleFunc("/logout/", handleLogout(store, secret)).Methods("POST")
}

type HandleLoginRequest struct {
	Email    string `json:"email,omitempty" example:"aboba@gmail.ru"` // either email or username should not be empty
	Username string `json:"username,omitempty" example:"aboba"`
	Password string `json:"password" example:"qwerty123456"`
	AppId    string `json:"app_id" example:"1"`
	ClientId string `json:"client_id" example:"1"`
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
// @Success 200 {object} jwt.TokenPair
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

		tokenPair, err := jwt.IssueTokenPair(store, user, req.AppId, req.ClientId, secret)
		if err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}

		respond(w, r, http.StatusOK, tokenPair)
	}
}

type HandleRefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" example:"07b7f432-7414-4340-890d-0376e46f1a00"`
}

// Register godoc
// @Summary Обновление JWT токена.
// @Description Принимает json с refresh-токеном.
// @Description Возвращает либо json с парой access-refresh токенами, либо ошибку.
// @Tags Auth
// @ID auth-refresh
// @Accept json
// @Produce json
// @Param Request body HandleRefreshTokenRequest true "json запроса:"
// @Success 200 {object} jwt.TokenPair
// @Failure 400	{object} HttpErrorResponse
// @Router /refresh/ [post]
func handleRefreshToken(store store.Store, secret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := &HandleRefreshTokenRequest{}

		if err := json.NewDecoder(r.Body).Decode(request); err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}

		tokenPair, err := jwt.IssueByRefreshToken(request.RefreshToken, store, secret)
		if err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}
		respond(w, r, http.StatusOK, tokenPair)

	}
}

// Register godoc
// @Summary Выход из аккаунта.
// @Description Принимает json с refresh-токеном.
// @Description Удаляет токен из базы данных. Либо ничего не возвращает, либо возвращает ошибку
// @Tags Auth
// @ID auth-logout
// @Accept json
// @Produce json
// @Param Request body HandleRefreshTokenRequest true "json запроса:"
// @Success 200 {string} OK
// @Failure 400	{object} HttpErrorResponse
// @Router /logout/ [post]
func handleLogout(store store.Store, secret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := &HandleRefreshTokenRequest{}

		if err := json.NewDecoder(r.Body).Decode(request); err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}

		if err := store.Token().Delete(request.RefreshToken); err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}

		respond(w, r, http.StatusOK, "OK!")

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
