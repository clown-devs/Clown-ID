package handlers

import (
	"clown-id/internal/models"
	"clown-id/internal/store"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterAuthHandlers(router *mux.Router, store store.Store, secret string) {
	router.HandleFunc("/login/", handleLogin(store, secret)).Methods("POST")
	router.HandleFunc("/register/", handleRegister(store)).Methods("POST")
}

// Login godoc
// @Summary Аутентификация пользователя по логину и паролю
// @Description Возвращает пару токенов - access и refresh токен или json с ошибкой
// @Tags Auth
// @ID auth-login
// @Produce json
// @Consume json
// @Success 200
// @Failure 404
// @Router /login [get]
func handleLogin(store store.Store, secret string) http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}

		user, err := store.User().FindByEmail(req.Email)
		if err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}

		//refreshToken := uuid.New().String()

		respond(w, r, http.StatusOK, user)
	}
}

func handleRegister(store store.Store) http.HandlerFunc {
	type request struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		request := &request{}

		if err := json.NewDecoder(r.Body).Decode(request); err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}

		u := &models.User{Username: request.Username, Email: request.Email, Password: request.Password}
		if err := store.User().Create(u); err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}

		u.Sanitize()
		respond(w, r, http.StatusOK, u)
	}

}
