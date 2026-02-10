package blogposts

import (
	"io"
	"io/fs"
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
	postContent, err := io.ReadAll(postMdFile)
	if err != nil {
		return Post{}, err
	}

	tags := []string{}
	return Post{string(postContent[7:]), "","", tags}, nil
}
