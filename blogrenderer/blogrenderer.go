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

func Convert(w io.Writer, p Post) error {
	template, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}

	if err := template.Execute(w, p); err != nil {
		return err
	}

	return nil
}
