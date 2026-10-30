package httpx

import (
	"encoding/json"
	"net/http"
)

type Code string

const (
	Code_invalid_id        Code = "invalid_id"
	Code_not_found         Code = "not_found"
	Code_internal_error    Code = "internal_error"
	Code_malformed_json    Code = "malformed_json"
	Code_validation_failed Code = "validation_failed"
	Code_unauthenticated   Code = "unauthenticated"
	Code_forbidden         Code = "forbidden"
	Code_conflict          Code = "conflict"
	Code_rate_limited      Code = "rate_limited"
)

type errorEnvelop struct {
	Error errorPayload `json:"error"`
}
type errorPayload struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
}

func Error(w http.ResponseWriter, status int, message string, code Code) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(errorEnvelop{Error: errorPayload{
		Code:    code,
		Message: message,
	}})
}
func ValidationError(w http.ResponseWriter, status int, message string, code Code, field string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(errorEnvelop{Error: errorPayload{
		Code:    code,
		Message: message,
		Field:   field,
	}})
}
