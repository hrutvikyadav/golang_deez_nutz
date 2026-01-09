package main

import (
	"fmt"
)

const englishGreetPrefix = "Hello, "
const spanishGreetPrefix = "Hola, "
const spanish = "spanish"

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	if lang == spanish {
		return spanishGreetPrefix + name
	}
	return englishGreetPrefix + name
}
func main() {
	fmt.Println(Hello("World", ""))
}
