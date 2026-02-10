package blogrenderer

import (
	"fmt"
	"io"
)

type Post struct {
	Title string
	Description string
	Tags []string
	Body string
}

func Convert(w io.Writer, p Post) error {
	_, err := fmt.Fprintf(w, "<h1>%s</h1>", p.Title)
	return err
}
