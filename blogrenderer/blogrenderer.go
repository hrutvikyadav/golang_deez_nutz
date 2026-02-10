package blogrenderer

import (
	"io"
	"text/template"
)

type Post struct {
	Title string
	Description string
	Tags []string
	Body string
}

const postTemplate = `<h1>{{.Title}}</h1><p>{{.Description}}</p>Tags: <ul>{{range .Tags}}<li>{{.}}</li>{{end}}</ul>`

func Convert(w io.Writer, p Post) error {
	template, err := template.New("blog").Parse(postTemplate)
	if err != nil {
		return err
	}

	if err := template.Execute(w, p); err != nil {
		return err
	}

	return nil
}
