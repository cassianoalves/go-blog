package presentation

import (
	"github.com/gorilla/mux"
)

/**
POST /posts: Create a new blog post.
GET /posts: Retrieve a list of all blog posts.
GET /posts/{id}: Retrieve a specific blog post by its ID.
PUT /posts/{id}: Update a specific blog post by its ID.
DELETE /posts/{id}: Delete a specific blog post by its ID.
POST /posts/{id}/comments: Add a comment to a specific blog post.
GET /posts/{id}/comments: Retrieve all comments for a specific blog post.
DELETE /posts/{postId}/comments/{commentId}: Delete a specific comment for a specific blog post.
*/

func DefineRoutes(router *mux.Router, pc *PostController) {

	router.HandleFunc("/posts", pc.CreatePost).Methods("POST")
	router.HandleFunc("/posts", GetAllPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", GetPost).Methods("GET")
	router.HandleFunc("/posts/{id}", UpdatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", DeletePost).Methods("DELETE")
	router.HandleFunc("/posts/{id}/comments", AddComment).Methods("POST")
	router.HandleFunc("/posts/{id}/comments", GetAllComments).Methods("GET")
	router.HandleFunc("/posts/{postId}/comments/{commentId}", DeleteComment).Methods("DELETE")

}
