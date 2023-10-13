package presentation

import (
	"github.com/stretchr/testify/assert"
	"goBlog/domain"
	"net/http"
	"net/http/httptest"
	"testing"
)

type spy struct {
	callCounter *int
}

type MockService struct {
	id  int
	err error
}

func (m MockService) Create(post domain.Post) (int, error) {
	return m.id, m.err
}

func TestGetPost(t *testing.T) {
	a := assert.New(t)
	mockService := &MockService{
		123,
		nil,
	}
	controller := &PostController{mockService}

	postIn := domain.Post{
		Title:   "Some Subject",
		Author:  "John Smith",
		Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
	}

	r, _ := http.NewRequest(http.MethodPost, "https://service.host/posts", postIn.ToJsonReader())
	w := httptest.NewRecorder()

	controller.CreatePost(w, r)

	a.Equal(0, w.Body.Len())
	a.Equal(http.StatusCreated, w.Code)
	a.Equal("https://service.host/posts/123", w.Header().Get("Location"))
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
