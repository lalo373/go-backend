package handlers

import (
	"backend/internal/httpx"
	"backend/middlewear"
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
	Price       uint64    `json:"ad-price"`
	City        string    `json:"ad-city"`
	CreatedAt   time.Time `json:"ad-created"`
}
type ListingConstructor struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewListingsHandler(db *sql.DB, logger *slog.Logger) *ListingConstructor {
	return &ListingConstructor{
		db:     db,
		logger: logger,
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
		slog.Error("query", "err", err)
		httpx.Error(w, http.StatusInternalServerError, "Internal error", httpx.Code_internal_error)

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
			httpx.Error(w, http.StatusInternalServerError, "Internal error", httpx.Code_internal_error)
			return
		}
		listings = append(listings, lists)
	}
	if err := rows.Err(); err != nil {
		slog.Error("rows.err", "err", err)
		httpx.Error(w, http.StatusInternalServerError, "Internal error", httpx.Code_internal_error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(listings)

}

func (lc ListingConstructor) DeleteListing(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := middlewear.RequestIDFromContext(ctx)
	id := r.PathValue("id")
	_, err := lc.db.ExecContext(ctx,
		`delete from listings where id = $1`, id) // using placeholder to prevent sql injection
	// defer db.Close()
	if err != nil {

		log.Printf("delete err : %v", err)
		lc.logger.Error("delete failed", "listing-id", id, "err", err, "remote-addr", r.RemoteAddr, "request-ID", requestID)
		// http.Error(w, "internal error", http.StatusInternalServerError)
		httpx.Error(w, http.StatusInternalServerError, "Internal error", httpx.Code_internal_error)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(id + "deleted sucessfully"))
}

func (lc ListingConstructor) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateListingReqReader
	reqID := middlewear.RequestIDFromContext(r.Context())
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		lc.logger.Error("failed to decode", "request_id", reqID, "error", err)
		httpx.Error(w, http.StatusBadRequest, "invalid Body", httpx.Code_malformed_json)
		return
	}
	if req.Title == "" || req.Description == "" || req.City == "" || req.Price == 0 {
		httpx.Error(w, http.StatusBadRequest,
			"title, description, price and city are required; send keys ad-title, ad-description, ad-price, ad-city",
			httpx.Code_validation_failed)
		return
	}
	row := lc.db.QueryRowContext(r.Context(), `
		INSERT INTO listings
		(title, description, price, city)
		VALUES ($1, $2, $3, $4) RETURNING id, title, created_at`, req.Title, req.Description, req.Price, req.City)
	var out CreateListingReqSender
	if err := row.Scan(&out.ID, &out.Title, &out.CreatedAt); err != nil {
		lc.logger.Error("failed to insert", "request_id", reqID, "error", err)
		httpx.Error(w, http.StatusInternalServerError, "Something Went Wrong", httpx.Code_internal_error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(out)

}
