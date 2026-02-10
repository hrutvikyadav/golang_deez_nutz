package blogposts

import (
	"io/fs"
	"testing/fstest"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func PostsFromFS(filesys fstest.MapFS) (posts []Post) {
	dir, _ := fs.ReadDir(filesys, ".")
	for range dir {
		posts = append(posts, Post{})
	}
	return
}
