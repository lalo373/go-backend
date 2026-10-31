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
type ValidateErrorMsg struct {
	Field string
	Msg   string
}

func (e ValidateErrorMsg) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Msg)
}

func (req CreateListingReqReader) ValidateTitle() error {
	if strings.TrimSpace(req.Title) == "" {
		return ValidateErrorMsg{Field: "title", Msg: "Empty Title, Include Title"}
	}
	return nil
}

func (req CreateListingReqReader) ValidateDescription() error {
	if strings.TrimSpace(req.Description) == "" {
		return ValidateErrorMsg{Field: "Description", Msg: "Empty Description, Include Description"}
	}
	return nil
}

func (req CreateListingReqReader) ValidatePrice() error {
	if req.Price <= 0 {
		return ValidateErrorMsg{Field: "Price", Msg: "Empty Price, Include Price"}
	}
	return nil
}

func (req CreateListingReqReader) ValidateCity() error {
	if strings.TrimSpace(req.City) == "" {
		return ValidateErrorMsg{Field: "title", Msg: "Empty Title, Include Title"}
	}
	return nil
}
