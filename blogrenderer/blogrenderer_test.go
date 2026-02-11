package blogrenderer_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/approvals/go-approval-tests"
	"github.com/approvals/go-approval-tests/reporters"
	"github.com/hrutvikyadav/blogrenderer"
)

func TestBlogRenderer(t *testing.T) {
	post := blogrenderer.Post{
		Title: "Hello once more",
		Description: "desc",
		Tags: []string{ "sometag", "othertag" },
		Body: `# Body
## some imp
lorem ipsum dolor emmet lorem ipsum dolor emmet lorem ipsum dolor emmet lorem ipsum dolor emmet lorem ipsum dolor emmet`,
	}

	postRenderer, err := blogrenderer.NewPostRenderer()

	if err != nil {
		t.Fatal(err)
	}

	t.Run("should generate html from Posts", func(t *testing.T) {

		buf := bytes.Buffer{}
		err := postRenderer.Convert(&buf, post)
		if err != nil {
			t.Error(err)
		}

		approvals.UseReporter(reporters.NewSystemoutReporter())
		approvals.VerifyString(t, buf.String())

	})

	t.Run("should render index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blogrenderer.Post{{Title: "blog1"}, {Title: "blog2"}, {Title: "Blog 4"}}

		err := postRenderer.RenderIndex(&buf, posts)
		if err != nil {
			t.Error(err)
		}

		want := `<ol><li><a href="/post/blog1">blog1</a></li><li><a href="/post/blog2">blog2</a></li><li><a href="/post/blog-4">Blog 4</a></li></ol>`
		got := buf.String()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})
}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)
	postRenderer, err := blogrenderer.NewPostRenderer()

	if err != nil {
		b.Fatal(err)
	}
	for b.Loop() {
		postRenderer.Convert(io.Discard, aPost)
	}
}
