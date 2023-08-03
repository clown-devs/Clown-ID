package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"clown-id/internal/store"

	"github.com/gorilla/mux"
)

// Middleware and handlers registering
func RegisterHandlers(router *mux.Router, store store.Store, secret string) {
	router.Use(commonMiddleware)
	RegisterAuthHandlers(router, store, secret)

	router.HandleFunc("/hello/", handleHello())
}

// handleHello godoc
// @Summary Тестовый эндпоинт для проверки работы автодокументации
// @Description Возвращает json "hello world"
// @Tags Test
// @ID hello
// @Produce json
// @Success 200
// @Failure 404
// @Router /hello [get]
func handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		respond(w, r, http.StatusOK, map[string]string{"hello": "world"})
	}
}

func respondError(w http.ResponseWriter, r *http.Request, code int, err error) {
	//TODO: Write logs
	respond(w, r, code, map[string]string{"error": err.Error(), "code": fmt.Sprint(code)})
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
