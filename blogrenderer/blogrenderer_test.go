package blogrenderer_test

import (
	"bytes"
	"testing"

	"github.com/hrutvikyadav/golang_deez_nutz/blogrenderer"
)

func TestBlogRenderer(t *testing.T) {
	t.Run("should generate html from Posts", func(t *testing.T) {
		post := blogrenderer.Post{
			Title: "Hello once more",
			Description: "desc",
			Tags: []string{ "sometag", "othertag" },
			Body: `# Body
## some imp
lorem ipsum dolor emmet lorem ipsum dolor emmet lorem ipsum dolor emmet lorem ipsum dolor emmet lorem ipsum dolor emmet`,
		}

		buf := bytes.Buffer{}
		err := blogrenderer.Convert(&buf, post)
		if err != nil {
			t.Error(err)
		}

		got := buf.String()
		want := `<h1>Hello once more</h1><p>desc</p>Tags: <ul><li>sometag</li><li>othertag</li></ul>`

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

	})
}
