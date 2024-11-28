package main

import (
	"fmt"
	hello "go-with-tests/hello"
	"go-with-tests/integers"
	"go-with-tests/iterations"
)

func main() {
	fmt.Println(hello.Hello("Jace", ""))    // From hello pkg
	fmt.Println(integers.Add(2, 3))         // From integers pkg
	fmt.Println(iterations.Repeat("jace ")) // From iterations pkg
}
