package main

import (
	"backend/internal/config"
	"backend/internal/db"
	"backend/internal/handlers"
	"backend/middlewear"
	"backend/static"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	slog_json_handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	})
	sloglogger := slog.New(slog_json_handler)
	slog.SetDefault(sloglogger)
	cfg := config.MustLoad()
	db, err := db.Connect(cfg.DBURL)
	if err != nil {
		log.Fatalf("main.db.connect: %v", err)
	}
	fmt.Println(("starting database"))
	fmt.Println(("starting server"))
	lh := handlers.NewListingsHandler(db, sloglogger)

	mux := http.NewServeMux()
	// Serve the embedded frontend for every path; more specific
	// patterns (healthz, listings) still take precedence.
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(static.Index)
	})
	mux.HandleFunc("GET /healthz", handlers.HealthzHandler)
	mux.HandleFunc("GET /listings", lh.ListingsHandler)
	mux.HandleFunc("DELETE /listings/{id}", lh.DeleteListing)

	middlewearHandler := middlewear.RequestID(mux)
	s := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      middlewearHandler,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 40,
		IdleTimeout:  time.Second * 60,
	}
	log.Printf("server is listening on %s", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("server failed; %v", err)
	}
}
