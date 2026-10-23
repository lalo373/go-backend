package middlewear

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type ctxkey int

const (
	requestIDKey ctxkey = iota
)
const (
	requestId = "X-Request-ID"
)

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get(requestId)
		if id == "" {
			id = uuid.NewString()
		}
		w.Header().Add(requestId, id)
		ctx := context.WithValue(r.Context(), "requestCtxId", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RequestIDFromContext(ctx context.Context) string {
	requestID := ctx.Value(requestIDKey).(string)
	return requestID
}
