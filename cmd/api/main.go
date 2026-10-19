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
	db, err := db.Connect(cfg.DBURL)
	if err != nil {
		log.Fatalf("main.db.connect: %v", err)
	}
	fmt.Println(("starting database"))
	fmt.Println(("starting server"))
	lh := handlers.NewListingsHandler(db)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.RootHandler)
	mux.HandleFunc("GET /healthz", handlers.HealthzHandler)
	mux.HandleFunc("GET /listings", lh.ListingsHandler)
	mux.HandleFunc("DELETE /listings/{id}", lh.DeleteListing)

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
