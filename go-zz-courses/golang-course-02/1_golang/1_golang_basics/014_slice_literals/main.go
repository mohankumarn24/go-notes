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
	printHeader("Slice Declaration")
	players := []string{"Sachin", "Dravid"}
	fmt.Println("Players:", players) // Players: [Sachin Dravid]

	// Access elements
	printHeader("Access Elements")
	// fmt.Println("Players:", players)              // Players: [Sachin Dravid]
	fmt.Printf("Players: %q\n", players)          // Players: ["Sachin" "Dravid"]
	fmt.Println("First:", players[0])             // First: Sachin
	fmt.Println("Last:", players[len(players)-1]) // Last: Dravid
	fmt.Println("Length:", len(players))          // Length: 2
	fmt.Println("Capacity:", cap(players))        // Capacity: 2

	// Update an element
	printHeader("Update Elements")
	players[0] = "Sachin Tendulkar"
	players[1] = "Rahul Dravud"
	fmt.Printf("Updated Players: %q\n", players) // Updated Players: ["Sachin Tendulkar" "Rahul Dravud"]

	// Arrays cannot use append().
	// Slices can grow using append().
	printHeader("Append an element")
	players = append(players, "Virat Kohli")

	fmt.Printf("After Append: %q\n", players) // After Append: ["Sachin Tendulkar" "Rahul Dravud" "Virat Kohli"]
	fmt.Println("Length:", len(players))      // Length: 3
	fmt.Println("Capacity:", cap(players))    // Capacity: 4

	// Declare an empty slice. Then append elements
	printHeader("Declare an empty slice. Then append elements")
	var nums []int

	// Append elements
	nums = append(nums, 10)
	nums = append(nums, 20, 30)
	fmt.Println("Nums:", nums) // Nums: [10 20 30]

	// Length and capacity
	fmt.Println("Length:", len(nums))   // Length: 3
	fmt.Println("Capacity:", cap(nums)) // Capacity: 3

	// Slice another slice
	// [start:end] -> includes start, excludes end
	printHeader("Slice Expression")
	sub := nums[1:3]               // Elements at index 1 and 2 -> [20 30]
	fmt.Println("Nums:", nums)     // Nums: [10 20 30]
	fmt.Println("Sub Slice:", sub) // Sub Slice: [20 30]
	/*
		fmt.Println(nums)     	// [10 20 30]
		fmt.Println(nums[0:3]) 	// [10 20 30]
		fmt.Println(nums[1:3]) 	// [20 30]
		fmt.Println(nums[2:3]) 	// [30]
		fmt.Println(nums[3:3]) 	// []


		0 <= start <= end <= len(slice)
		len(nums) = 3

		✓ nums[0:3]
		✓ nums[1:3]
		✓ nums[2:3]
		✓ nums[3:3]

		✗ nums[0:4]
		✗ nums[2:4]
		✗ nums[4:4]
	*/

	// Iterate over a slice
	printHeader("Iterate Over a Slice")
	for i, value := range nums {
		// fmt.Println("Index:", i, "Value:", value)
		fmt.Printf("Index: %d, Value: %d\n", i, value)
	}
	/*
		Index: 0, Value: 10
		Index: 1, Value: 20
		Index: 2, Value: 30
	*/

	// Create a Slice Using make()
	printHeader("Create a Slice Using make()")
	numbers := make([]int, 3, 5)

	fmt.Println("Numbers:", numbers)       // Numbers: [0 0 0]
	fmt.Println("Length:", len(numbers))   // Length: 3
	fmt.Println("Capacity:", cap(numbers)) // Capacity: 5

	// Copy a slice
	printHeader("Copy a Slice")
	src := []int{10, 20, 30}      // 'src' has length = 3, capacity = 3.
	dest := make([]int, len(src)) // Create destination slice. 'dest' is created with length = 3, capacity = 3
	// initial value: Zero values ([0 0 0])

	copy(dest, src)                   // Copy elements from 'src' to 'dest'
	src[0] = 99                       // Changing 'src' does not affect 'dest'
	fmt.Println("Source:", src)       // Source: [99 20 30]
	fmt.Println("Destination:", dest) // Destination: [10 20 30]
	/*
		Before modification:
		src  -> [10 20 30]
		dest -> [10 20 30]

		src[0] = 99

		After modification:
		src  -> [99 20 30]
		dest -> [10 20 30]

		'dest' doesn't change because it has its own underlying array
	*/

	// nil slice
	printHeader("Nil Slice")
	var values []int

	fmt.Println(values)                   // []
	fmt.Println(values == nil)            // true
	fmt.Println("Length:", len(values))   // Length: 0
	fmt.Println("Capacity:", cap(values)) // Capacity: 0

	// Append multiple values
	printHeader("Append Multiple Values")
	nums = append(nums, 40, 50, 60)

	fmt.Println("Nums:", nums)          // Nums: [10 20 30 40 50 60]
	fmt.Println("Length:", len(nums))   // Length: 6
	fmt.Println("Capacity:", cap(nums)) // Capacity: 6

	// Append one slice to another
	printHeader("Append One Slice to Another")
	a := []int{1, 2}
	b := []int{3, 4}

	a = append(a, b...)
	fmt.Println("Combined Slice:", a) // Combined Slice: [1 2 3 4]

	/*
		// Slice (dynamic)
		slice := []string{"Sachin", "Dravid"}

		// Array (fixed size)
		array := [...]string{"Sachin", "Dravid"}

		fmt.Printf("Slice Type: %T\n", slice)
		fmt.Printf("Array Type: %T\n", array)
	*/
}

func printHeader(title string) {
	fmt.Printf("\n========== %s ==========\n", title)
}
