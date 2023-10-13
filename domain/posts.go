package domain

import (
	"bytes"
	"encoding/json"
	"io"
)

type Post struct {
	Id        int    `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Content   string `json:"content,omitempty"`
	Author    string `json:"author,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
}

func (p Post) ToJsonReader() io.Reader {
	b, err := json.Marshal(p)
	if err != nil {
		return nil
	}

	return bytes.NewReader(b)
}

func (p Post) IsValid() bool {
	return true // TODO
}

type IPostService interface {
	Create(post Post) (int, error)
}

type PostService struct{}

func (p PostService) Create(post Post) (int, error) {
	//TODO implement me
	panic("implement me")
}
