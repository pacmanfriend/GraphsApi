package graph

import (
	"GraphsApi/internal/handlers"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc("/create", h.Create()).Methods(http.MethodPost)
}

func (h *handler) Create() http.HandlerFunc {
	var requestData struct {
		Matrix               [][]int `json:"matrix"`
		Start                int     `json:"start"`
		End                  int     `json:"end"`
		VertexCount          int     `json:"vertexCount"`
		ShowIntermediateInfo bool    `json:"showIntermediateInfo"`
	}

	return func(response http.ResponseWriter, request *http.Request) {
		err := json.NewDecoder(request.Body).Decode(&requestData)
		if err != nil {
			http.Error(response, err.Error(), http.StatusBadRequest)
		}

		response.WriteHeader(200)
		fmt.Println(requestData)
	}
}
