// It will be nice to reuse the GreetSpecification to test Greet.
// But we cannot just do that because the types dont match
//
package interactions_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"
    "github.com/hrutvikyadav/go-specs-greet/domain/interactions"
    "github.com/hrutvikyadav/go-specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	// here the specification wants something that it can call Greet on,
	// but instead we have given a function.
	// note that the signature matches, but the *shape* is incorrect
	// The solution to these kinds of problems is the ADAPTER Pattern
	// we will wrap our function in an adapter so that it satisfies the interface
	// specifications.GreetSpecification(t, go_specs_greet.Greet)
	specifications.GreetSpecification(t, specifications.GreetAdapter(interactions.Greet))

	t.Run("should default name to World", func(t *testing.T) {
		got := interactions.Greet("")
		want := "Hello, World"

		assert.Equal(t, want, got)
	})
}
