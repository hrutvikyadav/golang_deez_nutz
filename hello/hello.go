package main

import (
	"fmt"
)

const englishGreetPrefix = "Hello, "

func Hello(name string) string {
	return englishGreetPrefix + name
}
func main() {
	fmt.Println(Hello("World"))
}
