package model

import "html/template"

type Post struct {
	Status  string
	Title   string
	Date    string
	Summary string
	Body    template.HTML
	File    string
}
