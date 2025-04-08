package controller

import (
	"log/slog"
	"net/http"
)

func dummyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// do something before
		next.ServeHTTP(w, r)
		// do something after
	})
}

func TestMiddlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("========== middleware 1 ==========")

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func TestMiddlewareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("========== middleware 2 ==========")

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
