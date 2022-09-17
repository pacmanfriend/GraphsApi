package main

import (
	"GraphsApi/internal/graph"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
	"time"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Access-Control-Allow-Headers:", "*")
		response.Header().Set("Access-Control-Allow-Origin", "*")
		response.Header().Set("Access-Control-Allow-Methods", "*")

		if request.Method == "OPTIONS" {
			response.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(response, request)
		return
	})
}

func main() {
	router := mux.NewRouter()
	router.Use(CORS)

	graphHandler := graph.NewHandler()
	graphHandler.Register(router)

	startServer(router)
}

func startServer(router *mux.Router) {
	listener, err := net.Listen("tcp", ":5050")

	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Fatal(server.Serve(listener))
}
