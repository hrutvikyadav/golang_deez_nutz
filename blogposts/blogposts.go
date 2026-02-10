package blogposts

import (
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
	for range dir {
		posts = append(posts, Post{})
	}
	return posts, nil
}
