package specifications

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

type Greeter interface {
	Greet(name string) (string, error)
}

func GreetSpecification(t testing.TB, greeter Greeter) {
	got, err := greeter.Greet("Me")
	assert.NoError(t, err)
	assert.Equal(t, got, "Hello, Me")
}

type MeanGreeter interface {
	Curse(name string) (string, error)
}

func CurseSpecification(t testing.TB, mg MeanGreeter) {
	got, err := mg.Curse("you")
	assert.NoError(t, err)
	assert.Equal(t, got, "Go to Hell, you!")
}
