package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	//values
	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	//variables
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	//constants
	fmt.Println(s)

	const n = 500000000

	const d2 = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d2))

	fmt.Println(math.Sin(n))

	//string
	var i, j string = "Hello", "World"

	fmt.Print(i, "\n", j)
	fmt.Print(i, " ", j)

	var a1, b1 = 10, 20
	fmt.Print(a1, b1)

	var c1 string = "Hello"
	var d1 int = 15
	fmt.Printf("\n c has value: %v and type: %T\n", c1, c1)
	fmt.Printf("d has value: %v and type: %T", d1, d1)

	// implemetation of slice
	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)
}
