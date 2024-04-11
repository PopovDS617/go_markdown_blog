package post

import (
	"encoding/json"
	"fmt"
	"gomarkdownblog/internal/logger"
	"net/http"
)

func (i *Implementation) GetBySlug(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	logger.Info(slug)

	if slug == "" {
		w.WriteHeader(400)
		return
	}

	post, err := i.postService.GetPostBySlug((slug))

	if err != nil {

		fmt.Println(err)

		if err.Error() != "file not found" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(post)

}

func (i *Implementation) GetList(w http.ResponseWriter, r *http.Request) {

	posts, err := i.postService.GetAllPosts()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(posts)

}
