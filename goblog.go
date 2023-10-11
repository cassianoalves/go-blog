package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"goBlog/domain"
	"goBlog/presentation"
	"net/http"
)

func main() {
	setupAndStartHttpServer()
}

func setupAndStartHttpServer() {
	router := mux.NewRouter()

	pc := &presentation.PostController{Service: &domain.PostService{}}
	presentation.DefineRoutes(router, pc)
	presentation.SetupMiddleware(router)

	// Start the HTTP server and listen on port 8080
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}
}
