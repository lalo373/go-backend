package main

import (
	"backend/internal/config"
	"backend/internal/db"
	"backend/internal/handlers"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	cfg := config.MustLoad()
	_, err := db.Connect(cfg.DBURL)
	if err != nil {
		log.Fatalf("main.db.connect: %v", err)
	}
	fmt.Println(("starting database"))
	fmt.Println(("starting server"))
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", handlers.HealthzHandler)
	mux.HandleFunc("GET /listings", handlers.ListingsHandler)

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
