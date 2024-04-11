package repository

import "gomarkdownblog/internal/repository/post"

type PostRepository interface {
	GetAll() ([]post.PostData, error)
	GetOneBySlug(slug string) ([]byte, error)
}
