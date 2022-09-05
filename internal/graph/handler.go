package graph

import (
	"GraphsApi/internal/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc("/graph", h.Create()).Methods(http.MethodPost)
}

func (h *handler) Create() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {

	}
}
