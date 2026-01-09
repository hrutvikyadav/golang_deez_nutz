package main

import (
	"fmt"
)

const englishGreetPrefix = "Hello, "
const spanishGreetPrefix = "Hola, "
const frenchGreetPrefix = "Bonjour, "
const spanish = "spanish"
const french = "french"

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	if lang == spanish {
		return spanishGreetPrefix + name
	} else if lang == french {
		return frenchGreetPrefix + name
	}
	return englishGreetPrefix + name
}
func main() {
	fmt.Println(Hello("World", ""))
}
