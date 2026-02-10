package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
	"time"

	blogposts "github.com/hrutvikyadav/blogposts"
)

func TestBlogPosts(t *testing.T) {
	t.Run("read md files from fs", func(t *testing.T) {
		const (
			firstBody = `Title: Hello
Description: desc
Tags: sometag, othertag
---
Body Heading
lolololol
Body 1`

			secondBody = `Title: Foo
Description: desc
Tags: tagc, tagb
---
Body 2`
			thirdBody = `Title: Holaasdjfklajdkfjakljdfjadjfa
Description: desc
Tags: taga, tagb
---
Body 3`
		)

		fs := fstest.MapFS{
			"hello world.md": {
				Data:    []byte(firstBody),
				Mode:    0,
				ModTime: time.Time{},
				Sys:     nil,
			},
			"hello foo.md": {
				Data:    []byte(secondBody),
				Mode:    0,
				ModTime: time.Time{},
				Sys:     nil,
			},
			"hello again.md": {
				Data:    []byte(thirdBody),
				Mode:    0,
				ModTime: time.Time{},
				Sys:     nil,
			},
		}

		posts, err := blogposts.PostsFromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		got := posts[2]
		tags := []string{"sometag", "othertag"}
		want := blogposts.Post{"Hello", "desc", `Body Heading
lolololol
Body 1`,
			tags}

		assertPost(t, got, want)
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

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
