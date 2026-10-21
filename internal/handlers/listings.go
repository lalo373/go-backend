package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"time"
)

type listing struct {
	ID          string    `json:"id"`
	Title       string    `json:"ad-title"`
	Description string    `json:"ad-description"`
	Price       string    `json:"ad-price"`
	City        string    `json:"ad-city"`
	CreatedAt   time.Time `json:"ad-created"`
}
type ListingConstructor struct {
	db *sql.DB
}

func NewListingsHandler(db *sql.DB) *ListingConstructor {
	return &ListingConstructor{
		db: db,
	}
}

func (lc ListingConstructor) ListingsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	rows, err := lc.db.QueryContext(ctx,
		`select id, title, description, price, city, created_at
			from listings
			order by created_at desc
			limit 100;`)
	if err != nil {
		slog.Error("query: %v", err)
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
		slog.Error("rows.err: %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(listings)

}

func (lc ListingConstructor) DeleteListing(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	_, err := lc.db.ExecContext(ctx,
		`delete from listings where id = $1`, id) // using placeholder to prevent sql injection
	// defer db.Close()
	if err != nil {

		log.Printf("delete err : %v", err)
		slog.Error("delete failed", "listing-id", id, "err", err, "remote-addr", r.RemoteAddr)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(id + "deleted sucessfully"))
}
