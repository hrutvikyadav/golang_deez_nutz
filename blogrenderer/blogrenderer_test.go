package blogrenderer_test

import (
	"bytes"
	"testing"

	"github.com/approvals/go-approval-tests"
	"github.com/approvals/go-approval-tests/reporters"
	"github.com/hrutvikyadav/blogrenderer"
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

		approvals.UseReporter(reporters.NewSystemoutReporter())
		approvals.VerifyString(t, buf.String())

	})
}
