package blogrenderer

import (
	"embed"
	"io"
	"text/template"
)

type Post struct {
	Title string
	Description string
	Tags []string
	Body string
}

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	template, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{template}, nil
}

func (r *PostRenderer) Convert(w io.Writer, p Post) error {
	if err := r.templ.ExecuteTemplate(w, "blog.gohtml", p); err != nil {
		return err
	}

	return nil
}
