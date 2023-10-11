package presentation

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func SetupMiddleware(router *mux.Router) {
	// Apply a middleware to log each request
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("Request received: %s %s\n", r.Method, r.URL.Path)
			next.ServeHTTP(w, r)
		})
	})
}
