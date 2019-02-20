package main

import (
	"fmt"
)

const (
	englishPrefix = "Hello"
	spanishPrefix = "Hola"
	frenchPrefix  = "Bonjour"
)

func hello(n, l string) string {

	if n == "" {
		n = "World"
	}
	return greetingsPrefix(l) + ", " + n + "!"
}

func greetingsPrefix(l string) (prefix string) {
	switch l {
	case "French":
		prefix = frenchPrefix
	case "Spanish":
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}
	return
}
func main() {
	fmt.Println(hello("Bana", "Spanish"))
}
