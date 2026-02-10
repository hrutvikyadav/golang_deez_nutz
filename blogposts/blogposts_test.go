package blogposts_test

import (
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"
	"time"

	blogposts "github.com/hrutvikyadav/blogposts"
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

		posts, err := blogposts.PostsFromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fs) {
			t.Errorf("expected %d posts, got %d", len(fs), len(posts))
		}
	})

	t.Run("error in PostsFromFS returns nil posts and the error", func(t *testing.T){
		_, err := blogposts.PostsFromFS(StubFailingFS{})

		if err == nil {
			t.Fatal("expected error but did not get one")
		}
		t.Log(err)
	})
}

type StubFailingFS struct {}

func (s StubFailingFS) Open(filename string) (file fs.File, err error) {
	return nil, errors.New("oh no messed up")
}
