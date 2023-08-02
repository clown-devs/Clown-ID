package handlers

import (
	"clown-id/internal/store"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterAuthHandlers(router *mux.Router, store store.Store) {
	router.HandleFunc("/login/", handleLogin(store)).Methods("POST")
}

// Login godoc
// @Summary Авторизация пользователя.
// @Description Возвращает json(будет описан позже)
// @Tags Auth
// @ID auth-login
// @Produce json
// @Consume json
// @Success 200
// @Failure 404
// @Router /login [get]
func handleLogin(store store.Store) http.HandlerFunc {
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

		respond(w, r, http.StatusOK, user)
	}
}
