package main

import "fmt"

func main() {
	// Slices are the most common collection type in Go.
	// They are dynamic and can grow.
	//
	// []T    -> Slice (dynamic size)
	// [...]T -> Array (fixed size, compiler determines length)
	//
	// Syntax: []type{...}

	// Slice literal
	results := []string{"Sangam", "John"}

	fmt.Println("Results:", results)
	fmt.Println("First:", results[0])
	fmt.Println("Last:", results[len(results)-1])

	fmt.Println("Length:", len(results))
	fmt.Println("Capacity:", cap(results))

	// Update an element
	results[1] = "Priya"
	fmt.Println("Updated:", results)

	// Arrays cannot use append().
	// Slices can grow using append().
	results = append(results, "Alice")

	fmt.Println("After Append:", results)
	fmt.Println("Length:", len(results))
	fmt.Println("Capacity:", cap(results))

	// Declare an empty slice
	var nums []int

	// Append elements
	nums = append(nums, 10)
	nums = append(nums, 20, 30)

	fmt.Println("Nums:", nums)

	// Length and capacity
	fmt.Println("Length:", len(nums))
	fmt.Println("Capacity:", cap(nums))

	// Slice another slice
	// [start:end] -> includes start, excludes end
	sub := nums[1:3]
	fmt.Println("Sub Slice:", sub)

	// Iterate over a slice
	for i, value := range nums {
		fmt.Println("Index:", i, "Value:", value)
	}

	/*
		// Slice (dynamic)
		slice := []string{"Sangam", "John"}

		// Array (fixed size)
		array := [...]string{"Sangam", "John"}

		fmt.Printf("Slice Type: %T\n", slice)
		fmt.Printf("Array Type: %T\n", array)
	*/
}
