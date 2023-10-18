package presentation

import (
	"context"
	"fmt"
	"net/http"
)

type User struct {
	name string
}

func AuthenticationHandler() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user := getUser(r)
			if user == nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "user", user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func getUser(r *http.Request) *User {
	fmt.Printf("Basic Auth: %s\n", r.Header.Get("Authorization"))
	return &User{name: "lala"} // TODO
}
