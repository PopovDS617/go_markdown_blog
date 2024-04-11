package app

import (
	postImpl "gomarkdownblog/internal/api/post"
	"gomarkdownblog/internal/repository"
	postRepo "gomarkdownblog/internal/repository/post"
	"gomarkdownblog/internal/service"
	postSrv "gomarkdownblog/internal/service/post"
	"net/http"
)

type serviceProvider struct {
	postService    service.PostService
	postRepository repository.PostRepository
	Implementation *postImpl.Implementation
	HTTPRouter     *http.ServeMux
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) initPostService() service.PostService {
	if s.postService == nil {
		s.postService = postSrv.NewService(s.initPostRepository())
	}

	return s.postService
}

func (s *serviceProvider) initPostRepository() repository.PostRepository {
	if s.postRepository == nil {
		s.postRepository = postRepo.NewRepository()
	}

	return s.postRepository
}

func (s *serviceProvider) initImpl() *postImpl.Implementation {
	if s.Implementation == nil {
		s.Implementation = postImpl.NewImplementation(s.initPostService())
	}

	return s.Implementation
}

func (s *serviceProvider) initHTTPRouter() *http.ServeMux {

	s.initImpl()

	if s.HTTPRouter == nil {
		mux := http.NewServeMux()

		mux.HandleFunc("GET /posts/{slug}/", s.Implementation.GetBySlug)
		mux.HandleFunc("GET /posts/", s.Implementation.GetList)

		s.HTTPRouter = mux

	}

	return s.HTTPRouter
}
