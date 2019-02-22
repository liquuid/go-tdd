package di

import (
	"bytes"
	"fmt"
)

func Greet(writer *bytes.Buffer, s string) {
	fmt.Fprintf(writer, "Ola, %s", s)
}