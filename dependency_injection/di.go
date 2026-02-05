package main

import (
	"bytes"
	"fmt"
)

func Greet(buf *bytes.Buffer, s string) {
	// NOTE: `fmt.Printf("Hello, %s", s)` this is the simpler version but it is not testable
	// Printf prints to standard output which cannot be tested using a test framework.
	// So we use dependency injection to inject the dependency of Printf (which is identified by digging ballz deep),
	// this way we pass in (inject the dependency to have more control over printing in a way we can test)
	fmt.Fprintf(buf, "Hello, %s", s)
}
