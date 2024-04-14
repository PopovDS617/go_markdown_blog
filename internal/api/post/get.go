package post

import (
	"gomarkdownblog/internal/logger"
	"net/http"
	"text/template"
)

func (i *Implementation) GetBySlug(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	if slug == "" {
		w.WriteHeader(400)
		return
	}

	post, err := i.postService.GetPostBySlug((slug))

	if err != nil {

		logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	t := template.New("post.html")
	t, err = t.ParseFiles("templates/post.html")

	if err != nil {
		logger.Error(err.Error())

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := t.Execute(w, post); err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}

func (i *Implementation) GetList(w http.ResponseWriter, r *http.Request) {

	posts, err := i.postService.GetAllPosts()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	t := template.New("index.html")
	t, err = t.ParseFiles("templates/index.html")
	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := t.Execute(w, posts); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
