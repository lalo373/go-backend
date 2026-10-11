package main

import (
	"backend/internal/config"
	"backend/internal/handlers"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	cfg := config.MustLoad()
	mux := http.NewServeMux()
	fmt.Println(("starting server"))
	mux.HandleFunc("GET /healthz", handlers.Handler)

	s := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 40,
		IdleTimeout:  time.Second * 60,
	}
	log.Printf("server is listening on %s", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("server failed; %v", err)
	}
}
