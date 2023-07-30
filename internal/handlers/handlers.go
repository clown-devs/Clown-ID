package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"clown-id/internal/store"

	"github.com/gorilla/mux"
)

// Middleware and handlers registering
func RegisterHandlers(router *mux.Router, store store.Store) {
	router.Use(commonMiddleware)
}

func respondError(w http.ResponseWriter, r *http.Request, code int, err error) {
	respond(w, r, code, map[string]string{"error": err.Error(), "code": fmt.Sprint(code)})
}

func respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			respondError(w, r, code, err)
		}
	}
}
