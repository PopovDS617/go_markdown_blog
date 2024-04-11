package post

import "gomarkdownblog/internal/service"

type Implementation struct {
	postService service.PostService
}

func NewImplementation(postService service.PostService) *Implementation {
	return &Implementation{
		postService,
	}
}
