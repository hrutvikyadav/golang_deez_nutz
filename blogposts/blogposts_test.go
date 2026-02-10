package blogposts_test

import (
	blogposts "github.com/hrutvikyadav/blogposts"
	"testing"
	"testing/fstest"
	"time"
)

func TestBlogPosts(t *testing.T) {
	t.Run("read md files from fs", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md": {
				Data:    []byte("hi"),
				Mode:    0,
				ModTime: time.Time{},
				Sys:     nil,
			},
			"hello again.md": {
				Data:    []byte("hola"),
				Mode:    0,
				ModTime: time.Time{},
				Sys:     nil,
			},
		}

		posts := blogposts.PostsFromFS(fs)

		if len(posts) != len(fs) {
			t.Errorf("expected %d posts, got %d", len(fs), len(posts))
		}
	})
}
