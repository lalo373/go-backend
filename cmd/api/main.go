package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"OK"}`))
	})

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
