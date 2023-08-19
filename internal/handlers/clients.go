package handlers

import (
	"clown-id/internal/store"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterClientHandlers(router *mux.Router, store store.Store, secret string) {
	router.HandleFunc("/clients/", handleGetClients(store)).Methods("GET")
	router.HandleFunc("/apps/", handleGetApps(store)).Methods("GET")
}

// Register godoc
// @Summary Получение списка клиентов.
// @Description Возвращает json со списком возможных клиентов.
// @Tags Clients
// @ID auth-clients
// @Produce json
// @Success 200 {array} models.Client
// @Router /clients [get]
func handleGetClients(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clients, err := store.Client().AllClients()
		if err != nil {
			respondError(w, r, http.StatusInternalServerError, err)
			return
		}

		respond(w, r, http.StatusOK, clients)
	}
}

// Register godoc
// @Summary Получение списка приложений.
// @Description Возвращает json со списком приложений зарегистрированных в этом сервисе.
// @Tags Clients
// @ID auth-apps
// @Produce json
// @Success 200 {array} models.Application
// @Router /apps [get]
func handleGetApps(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apps, err := store.Client().AllApps()
		if err != nil {
			respondError(w, r, http.StatusInternalServerError, err)
			return
		}

		respond(w, r, http.StatusOK, apps)
	}
}
