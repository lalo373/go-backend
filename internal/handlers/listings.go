package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type listing struct {
	ID          string
	Title       string
	Description string
	Price       string
	City        string
	CreatedAt   time.Time
}

func ListingsHandler(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		rows, err := db.Query(
			`select id, title, description, price, city, created_at
			from listings
			order by created_at desc
			limit 100;`)
		if err != nil {
			log.Printf("query: %v", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		listings := []listing{}
		for rows.Next() {
			var lists listing
			if err := rows.Scan(
				&lists.ID,
				&lists.Title,
				&lists.Description,
				&lists.Price,
				&lists.City,
				&lists.CreatedAt,
			); err != nil {
				log.Printf("rows scan error: %v", err)
				http.Error(w, "internal error", http.StatusInternalServerError)
				return
			}
			listings = append(listings, lists)
		}
		if err := rows.Err(); err != nil {
			log.Printf("rows.err: %v", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(listings)
	}
}
