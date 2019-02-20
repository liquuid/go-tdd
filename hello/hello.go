package main

import (
	"fmt"
)

const englishPrefix = "Hello"

func hello(n string) string {
	if n == "" {
		n = "World"
	}
	return englishPrefix + ", " + n + "!"
}

func main() {
	fmt.Println(hello("Bana"))
}
