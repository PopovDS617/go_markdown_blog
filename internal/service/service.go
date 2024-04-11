package service

import "gomarkdownblog/internal/model"

type PostService interface {
	GetPostBySlug(slug string) (model.Post, error)
	GetAllPosts() ([]model.Post, error)
}
