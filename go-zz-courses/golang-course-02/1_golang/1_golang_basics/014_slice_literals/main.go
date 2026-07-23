package main

import "fmt"

func main() {
	// Slices are the most common collection type in Go.
	// They are dynamic and can grow.
	// Syntax: []type{...}

	results := []string{"Sangam", "John"}
	fmt.Println("Results:", results)
	fmt.Println("First:", results[0])
	fmt.Println("Last:", results[len(results)-1])

	// Update an element
	results[1] = "Priya"
	fmt.Println("Updated:", results)

	// Append elements
	var nums []int

	nums = append(nums, 10)
	nums = append(nums, 20, 30)

	fmt.Println("Nums:", nums)

	// Length and capacity
	fmt.Println("Length:", len(nums))
	fmt.Println("Capacity:", cap(nums))

	// Slice another slice
	sub := nums[1:3]
	fmt.Println("Sub Slice:", sub)

	// Iterate over a slice
	for i, value := range nums {
		fmt.Println(i, value)
	}
}
