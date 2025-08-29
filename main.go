package main

import (
	"fmt"
	"health-service/internal/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", handler.PingHandler)
	mux.HandleFunc("GET /health", handler.HealthHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("🚀 Server started on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
