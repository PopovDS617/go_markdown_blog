package post

import (
	"gomarkdownblog/internal/model"
	"gomarkdownblog/internal/repository"
	"html/template"
	"strings"

	"github.com/russross/blackfriday"
)

type Service struct {
	postRepository repository.PostRepository
}

func NewService(repo repository.PostRepository) Service {

	return Service{
		postRepository: repo,
	}
}

func (s Service) GetPostBySlug(slug string) (model.Post, error) {

	var post model.Post

	fileread, err := s.postRepository.GetOneBySlug(slug)

	if err != nil {
		return post, err
	}

	lines := strings.Split(string(fileread), "\n")

	status := string(lines[0])
	title := string(lines[1])
	date := string(lines[2])
	summary := string(lines[3])
	body := strings.Join(lines[4:], "\n")
	htmlBody := template.HTML(blackfriday.MarkdownCommon([]byte(body)))
	post = model.Post{Status: status, Title: title, Date: date, Summary: summary, Body: htmlBody,

		File: slug,
	}

	return post, nil

}

func (s Service) GetAllPosts() ([]model.Post, error) {
	postDataSlice, err := s.postRepository.GetAll()

	if err != nil {
		return nil, err
	}

	posts := make([]model.Post, 0, len(postDataSlice))

	for _, v := range postDataSlice {

		lines := strings.Split(string(v.FileRead), "\n")
		status := string(lines[0])
		title := string(lines[1])
		date := string(lines[2])
		summary := string(lines[3])
		body := strings.Join(lines[4:], "\n")
		htmlBody := template.HTML(blackfriday.MarkdownCommon([]byte(body)))

		post := model.Post{
			Status: status, Title: title, Date: date, Summary: summary, Body: htmlBody, File: v.File,
		}

		posts = append(posts, post)

	}

	return posts, err
}
