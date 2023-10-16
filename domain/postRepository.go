package domain

type IPostRepository interface {
	Create(post Post, user string) (id string, err error)
}
