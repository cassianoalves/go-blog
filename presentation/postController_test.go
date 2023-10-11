package presentation

import (
	"github.com/stretchr/testify/assert"
	"goBlog/domain"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockService struct {
	validate bool
}

func (m MockService) Validate() bool {
	return m.validate
}

func TestGetPost(t *testing.T) {
	a := assert.New(t)
	mockService := &MockService{}
	controller := &PostController{mockService}

	postIn := domain.Post{
		Title:   "Some Subject",
		Author:  "John Smith",
		Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
	}

	r, _ := http.NewRequest(http.MethodPost, "/posts", postIn.ToJsonReader())
	w := httptest.NewRecorder()

	controller.CreatePost(w, r)

	a.Equal(w.Body.Len(), 0)
	a.Equal(w.Code, http.StatusCreated)

}

func TestGetPost1(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetPost(tt.args.w, tt.args.r)
		})
	}
}

/*

func TestAddComment(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddComment(tt.args.w, tt.args.r)
		})
	}
}

func TestDeleteComment(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteComment(tt.args.w, tt.args.r)
		})
	}
}

func TestDeletePost(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeletePost(tt.args.w, tt.args.r)
		})
	}
}

func TestGetAllComments(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetAllComments(tt.args.w, tt.args.r)
		})
	}
}

func TestGetAllPosts(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetAllPosts(tt.args.w, tt.args.r)
		})
	}
}



func TestPostController_CreatePost(t *testing.T) {
	type fields struct {
		service domain.PostService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := &PostController{
				service: tt.fields.service,
			}
			pc.CreatePost(tt.args.w, tt.args.r)
		})
	}
}

func TestUpdatePost(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdatePost(tt.args.w, tt.args.r)
		})
	}
}



*/
