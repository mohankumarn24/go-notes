package main

import "fmt"

func main() {
	// ==============================
	// Slice Creation using make()
	// make([]T, length, capacity)
	// ==============================

	scores := make([]int, 0, 5)

	fmt.Println("Initial Slice")
	fmt.Println("scores  :", scores)
	fmt.Println("length  :", len(scores))
	fmt.Println("capacity:", cap(scores))
	fmt.Println()

	// ==============================
	// Append one element
	// ==============================

	scores = append(scores, 100)

	fmt.Println("After append(100)")
	fmt.Println(scores)
	fmt.Println("length  :", len(scores))
	fmt.Println("capacity:", cap(scores))
	fmt.Println()

	// ==============================
	// Append multiple elements
	// ==============================

	scores = append(scores, 200, 300)

	fmt.Println("After append(200, 300)")
	fmt.Println(scores)
	fmt.Println("length  :", len(scores))
	fmt.Println("capacity:", cap(scores))
	fmt.Println()

	scores = append(scores, 45, 55)

	fmt.Println("After append(45, 55)")
	fmt.Println(scores)
	fmt.Println("length  :", len(scores))
	fmt.Println("capacity:", cap(scores))
	fmt.Println()

	// ==========================================
	// Exceeding capacity
	// Go automatically allocates a larger
	// backing array when needed.
	// (Growth strategy is implementation-dependent)
	// ==========================================

	scores = append(scores, 60)

	fmt.Println("After append(60)")
	fmt.Println(scores)
	fmt.Println("length  :", len(scores))
	fmt.Println("capacity:", cap(scores))
	fmt.Println()

	// ==============================
	// Access elements
	// ==============================

	fmt.Println("First :", scores[0])
	fmt.Println("Last  :", scores[len(scores)-1])
	fmt.Println()

	// ==============================
	// Slice operations
	// ==============================

	fmt.Println(scores)                       // [100 200 300 45 55 60]
	fmt.Println("scores[:3] =", scores[:3])   // [100 200 300]
	fmt.Println("scores[2:] =", scores[2:])   // [300 45 55 60]
	fmt.Println("scores[1:4] =", scores[1:4]) // [200 300 45]
	fmt.Println()

	// ==============================
	// Append one slice into another
	// ==============================

	todos := []string{
		"Do YouTube",
		"Workout Everyday",
	}

	more := []string{
		"Learn Golang",
		"Read Clean Code",
	}

	todos = append(todos, more...)

	fmt.Println("Todos:", todos)
}
