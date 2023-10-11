package presentation

import (
	"fmt"
	"goBlog/domain"
	"net/http"
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

type PostController struct {
	Service domain.IPostService
}

func (pc *PostController) CreatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreatePost")
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetAllPosts")
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetPost")
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdatePost")
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeletePost")
}

func AddComment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AddComment")
}

func GetAllComments(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetAllComments")
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetAllComments")
}
