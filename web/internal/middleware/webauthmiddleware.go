package middleware

import (
	"context"
	"net/http"
)

type WebAuthMiddleware struct {
}

func NewWebAuthMiddleware() *WebAuthMiddleware {
	return &WebAuthMiddleware{}
}

func (m *WebAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		val := r.Header.Get("User-Agent")
		reqCtx := r.Context()
		ctx := context.WithValue(reqCtx, "User-Agent", val)
		newReq := r.WithContext(ctx)
		next(w, newReq)
	}
}
