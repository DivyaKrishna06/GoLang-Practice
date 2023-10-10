package main

import (
	"fmt"
)

func main() {
	var i, j string = "Hello", "World"

	fmt.Print(i, "\n", j)
	fmt.Print(i, " ", j)

	var a, b = 10, 20
	fmt.Print(a, b)

	var c string = "Hello"
	var d int = 15
	fmt.Printf("\n c has value: %v and type: %T\n", c, c)
	fmt.Printf("d has value: %v and type: %T", d, d)
}
