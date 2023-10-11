package main

import "fmt"

func main() {

	//array initialization
	arr1 := [5]int{}              //not initialized
	arr2 := [5]int{1, 2}          //partially initialized
	arr3 := [5]int{1, 2, 3, 4, 5} //fully initialized

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	//length of array
	arr1 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	arr2 := [...]int{1, 2, 3, 4, 5, 6}

	fmt.Println(len(arr1))
	fmt.Println(len(arr2))

	//car array
	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Print(cars)

	//2D array
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
