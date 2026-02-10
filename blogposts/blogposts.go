package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func PostsFromFS(filesys fs.FS) (posts []Post, err error) {
	dir, err := fs.ReadDir(filesys, ".")
	if err != nil {
		return nil, err
	}
	for _, d := range dir {
		post, err := getPost(filesys, d.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fsys fs.FS, d string) (Post, error) {
	postMdFile, err := fsys.Open(d)
	if err != nil {
		return Post{}, err
	}
	defer postMdFile.Close()

	return newPost(postMdFile)
}
func newPost(postMdFile io.Reader) (Post, error) {
	const (
		titleOffset = 7
		descOffset = 13
		tagOffset = 6
	)
	scanner := bufio.NewScanner(postMdFile)

	readline := func() string {
		scanner.Scan()
		return scanner.Text()
	}

	title := readline()[titleOffset:]
	desc := readline()[descOffset:]
	tags := strings.Split(readline()[tagOffset:], ", ")
	body := readBody(scanner)

	return Post{title, desc, body , tags}, nil
}
func readBody(scanner *bufio.Scanner) string {
	scanner.Scan()

	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n")
}
