package main

import (
	"fmt"
)

const englishGreetPrefix = "Hello, "
const spanishGreetPrefix = "Hola, "
const frenchGreetPrefix = "Bonjour, "
const spanish = "spanish"
const french = "french"

func getPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishGreetPrefix
	case french:
		prefix = frenchGreetPrefix
	default :
		prefix = englishGreetPrefix
	}
	return
}
func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	return getPrefix(lang) + name
}
func main() {
	fmt.Println(Hello("World", ""))
}
