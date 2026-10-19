package handlers

import (
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Welcome!!!, for checking the health visit /healthz, for checking the listings posted sofar visit /listings"))
}
