package handlers

import "time"

type CreateListingReqReader struct {
	Title       string `json:"ad-title"`
	Description string `json:"ad-description"`
	Price       uint64 `json:"ad-price"`
	City        string `json:"ad-city"`
}
type CreateListingReqSender struct {
	ID        string    `json:"id"`
	Title     string    `json:"ad-title"`
	CreatedAt time.Time `json:"ad-created"`
}
