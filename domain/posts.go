package domain

import (
	"bytes"
	"encoding/json"
	"io"
)

type Post struct {
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

type IPostService interface {
	Validate() bool
}

type PostService struct{}

func (p PostService) Validate() bool {
	//TODO implement me
	panic("implement me")
}
