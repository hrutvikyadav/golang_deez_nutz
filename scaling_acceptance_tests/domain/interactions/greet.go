package interactions

import "fmt"

func Greet(name string) string {
	if name == "" { name = "World" }
	return fmt.Sprintf("Hello, %s", name)
}

func Curse(name string) string {
	return fmt.Sprintf("Go to Hell, %s!", name)
}
