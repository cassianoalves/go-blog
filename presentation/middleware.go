package presentation

import (
	"encoding/json"
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

type BlogError struct {
	ErrorMsg string `json:"error_msg,omitempty"`
}

func BuildError(err error, w http.ResponseWriter, httpStatus int) {
	w.WriteHeader(httpStatus)
	jsonBytes, _ := json.Marshal(BlogError{err.Error()})
	_, err = w.Write(jsonBytes)
	if err != nil {
		fmt.Printf("Error writing response: %s\n", err.Error())
	}
}

func BuildCreateResponse(w http.ResponseWriter, r *http.Request, id int) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Location",
		fmt.Sprintf("%s/%d", r.URL, id),
	)
}
