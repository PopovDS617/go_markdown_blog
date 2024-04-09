package post

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type repo struct{}

func NewRepository() repo {
	return repo{}
}

func (r repo) GetAll() ([]PostData, error) {

	var result []PostData

	files, _ := filepath.Glob("assets/posts/*")
	for _, f := range files {
		var postData PostData

		file := strings.Replace(f, "posts/", "", -1)
		file = strings.Replace(file, ".md", "", -1)
		fileread, err := os.ReadFile(f)

		if err != nil {
			return result, err
		}

		postData.File = file
		postData.FileRead = fileread

		result = append(result, postData)

	}

	// var result []model.Post

	// for _, v := range posts {

	// 	post := FromPostRepoModelToModel(v)

	// 	result = append(result, post)
	// }

	// return result, nil

	return result, nil
}

func (r repo) GetOneBySlug(slug string) ([]byte, error) {

	var fileread []byte

	f := "assets/posts/" + slug + ".md"

	fileread, err := os.ReadFile(f)

	if err != nil {
		return fileread, err
	}

	if len(fileread) == 0 {
		return fileread, fmt.Errorf("file is empty")
	}

	return fileread, nil

}
