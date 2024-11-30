package main

import (
	"fmt"
	"go-with-tests/hello"
	"go-with-tests/integers"
	"go-with-tests/iterations"
	"go-with-tests/slices"
)

func main() {
	fmt.Println(hello.Hello("Jace", "")) // From hello pkg

	fmt.Println(integers.Add(2, 3)) // From integers pkg

	fmt.Println(iterations.Repeat("jace ")) // From iterations pkg

	fmt.Println(slices.MyArray1)                        // From slices pkg
	fmt.Println(slices.MyArray2)                        // From slices pkg
	fmt.Println(slices.SumArray([5]int{1, 3, 5, 7, 9})) // From slices pkg

	fmt.Println(slices.MySlice1)                    // From slices pkg
	fmt.Println(slices.MySlice2)                    // From slices pkg
	fmt.Println(slices.SumSlice([]int{2, 4, 6, 8})) // From slices pkg
	fmt.Println()

}
