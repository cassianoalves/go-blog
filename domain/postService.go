package domain

type IPostService interface {
	Create(post Post, userName string) (string, error)
}

type PostService struct {
	Repository IPostRepository
}

func (p PostService) Create(post Post, userName string) (string, error) {
	return p.Repository.Create(post, userName)
}
