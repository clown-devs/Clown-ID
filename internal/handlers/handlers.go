package handlers

import (
	"encoding/json"
	"net/http"

	"clown-id/internal/store"

	"github.com/gorilla/mux"
)

// Middleware and handlers registering
func RegisterHandlers(router *mux.Router, store store.Store, secret string) {
	router.Use(commonMiddleware)
	RegisterAuthHandlers(router, store, secret)
	RegisterClientHandlers(router, store, secret)

}

type HttpErrorResponse struct {
	Code  int    `json:"code" example:"400"`
	Error string `json:"error" example:"something went wrong..."`
}

func respondError(w http.ResponseWriter, r *http.Request, code int, err error) {
	//TODO: Write logs
	respond(w, r, code, HttpErrorResponse{code, err.Error()})
}

func respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	//TODO: Write logs
	w.WriteHeader(code)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			respondError(w, r, code, err)
		}
	}
}
