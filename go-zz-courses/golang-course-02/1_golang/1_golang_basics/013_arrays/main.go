package main

import "fmt"

func main() {
	// Array declaration and initialization
	var nums [3]int = [3]int{1, 2, 3}
	fmt.Println("Nums:", nums)

	// Arrays have a fixed size and cannot grow.
	var marks [3]int
	marks[0] = 10
	marks[1] = 20
	marks[2] = 50

	fmt.Println("Marks:", marks)
	fmt.Println("Third Mark:", marks[2])

	// Array literal
	res := [5]int{2, 3, 4, 5, 6}
	fmt.Println("Length:", len(res))

	// Let Go infer the array length
	colors := [...]string{"Red", "Green", "Blue"}
	fmt.Println("Colors:", colors)

	// Arrays are initialized with zero values
	var values [3]int
	fmt.Println("Zero Values:", values)

	// Arrays can be compared
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	fmt.Println("Arrays Equal:", a == b)

	// Iterate over an array
	for i, value := range res {
		fmt.Println(i, value)
	}
}
