package handlers

import (
	"fmt"
	"strings"
	"time"
)

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
type ValidateTitleErrorMsg struct {
	Field string
	Msg   string
}

func (e ValidateTitleErrorMsg) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Msg)
}

func (req CreateListingReqReader) ValidateTitle() error {
	if strings.TrimSpace(req.Title) == "" {
		return ValidateTitleErrorMsg{Field: "title", Msg: "Empty Title, Include Title"}
	}
	return nil
}
