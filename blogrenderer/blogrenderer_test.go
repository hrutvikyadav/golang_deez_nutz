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
