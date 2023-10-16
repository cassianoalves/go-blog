package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"goBlog/domain"
	"goBlog/presentation"
	"goBlog/repository"
	"net/http"
	"os"
)

func main() {
	ctx := context.Background()

	setupAndStartHttpServer(ctx)
}

func setupAndStartHttpServer(ctx context.Context) {
	router := mux.NewRouter()

	// TODO - refactor injections
	firestoreClient, fsErr := createFirestoreClient(ctx)
	if fsErr != nil {
		fmt.Println(fsErr)
		os.Exit(1)
	}
	pc := &presentation.PostController{
		Service: &domain.PostService{
			Repository: &repository.PostFirebaseDao{
				FirestoreClient:   firestoreClient,
				CollectionsPrefix: "PRD",
				Context:           ctx,
			},
		},
	}
	presentation.DefineRoutes(router, pc)
	presentation.SetupMiddleware(router)

	// Start the HTTP server and listen on port 8080
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}
}

func createFirestoreClient(ctx context.Context) (*firestore.Client, error) {
	// Use gcloud auth application-default login to set login to app
	projectID := "go-blog-401921" // TODO - external config
	return firestore.NewClient(ctx, projectID)
}
