package main

import "fmt"

func main() {
	// Array Declaration
	printHeader("Array Declaration")
	var nums [3]int = [3]int{1, 2, 3}
	fmt.Println("Nums:", nums) // Nums: [1 2 3]

	// Arrays have a fixed size and cannot grow
	printHeader("Arrays have a fixed size and cannot grow")
	var marks [3]int
	marks[0] = 10
	marks[1] = 20
	marks[2] = 50
	fmt.Println("Marks:", marks)         // Marks: [10 20 50]
	fmt.Println("Third Mark:", marks[2]) // Third Mark: 50

	// Array literal
	printHeader("Array Literal")
	res := [5]int{2, 3, 4, 5, 6}
	fmt.Println("Length:", len(res)) // Length: 5

	// Iterate over an array
	printHeader("Iterate over an array")
	for i, value := range res {
		output := fmt.Sprintf("Index: %d, Value: %d", i, value)
		fmt.Println(output)
	}

	// Let Go infer the array length
	printHeader("Let Go infer the array length")
	colors := [...]string{"Red", "Green", "Blue"}
	fmt.Println("Colors:", colors) // Colors: [Red Green Blue]

	// Arrays are initialized with zero values
	printHeader("Arrays are initialized with zero values")
	var values [3]int
	fmt.Println("Zero Values:", values) // Zero Values: [0 0 0]

	// Arrays can be compared
	printHeader("Arrays can be compared")
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	fmt.Println("Arrays Equal:", a == b) // Arrays Equal: true

	// 2D matrix
	// 3 rows, 4 columns
	printHeader("2D Matrix")
	matrix := [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(matrix)       // [[1 2 3 4] [5 6 7 8] [9 10 11 12]]
	fmt.Println(matrix[0][0]) // 1
	fmt.Println(matrix[1][2]) // 7
	fmt.Println(matrix[2][3]) // 12

	// Modify an Element
	printHeader("Modify an Element")
	matrix[1][2] = 100
	fmt.Println(matrix) // [[1 2 3 4] [5 6 100 8] [9 10 11 12]]

	// Iterate Using nested loops
	printHeader("Iterate Using nested loops")
	for i := range matrix {
		for j := range matrix[i] {
			// output := fmt.Sprintf("matrix[%d][%d] = %d", i, j, matrix[i][j])
			// fmt.Println(output)

			fmt.Printf("matrix[%d][%d] = %d\n", i, j, matrix[i][j])

		}
	}

	// Print row by row
	printHeader("Print row by row")
	fmt.Println("Matrix:")
	for _, row := range matrix {
		fmt.Println(row)
	}

	// Show len() for a 2D array
	printHeader("Show len() for a 2D array")
	fmt.Println("Rows:", len(matrix))
	fmt.Println("Columns:", len(matrix[0]))

	// Show zero values for a 2D array
	printHeader("Show zero values for a 2D array")
	var grid [2][3]int
	fmt.Println("Zero Matrix:")
	for _, row := range grid {
		fmt.Println(row)
	}

	// sum of all elements
	printHeader("Sum of All Elements")
	sum := 0
	for _, row := range matrix {
		for _, value := range row {
			sum += value
		}
	}
	fmt.Println("Sum:", sum)

	// Iterate using value instead of index
	printHeader("Iterate Using Values")
	for _, row := range matrix {
		for _, value := range row {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}

	// Modify every element
	printHeader("Modify every element")
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] *= 2
		}
	}

	fmt.Println("After Doubling:")
	for _, row := range matrix {
		fmt.Println(row)
	}

	// Compare two matrices
	printHeader("Compare two matrices")
	m1 := [2][2]int{
		{1, 2},
		{3, 4},
	}

	m2 := [2][2]int{
		{1, 2},
		{3, 4},
	}

	fmt.Println("Matrices Equal:", m1 == m2)
}

func printHeader(title string) {
	fmt.Printf("\n========== %s ==========\n", title)
}

/*
========== Array Declaration ==========
Nums: [1 2 3]

========== Arrays have a fixed size and cannot grow ==========
Marks: [10 20 50]
Third Mark: 50

========== Array Literal ==========
Length: 5

========== Iterate over an array ==========
Index: 0, Value: 2
Index: 1, Value: 3
Index: 2, Value: 4
Index: 3, Value: 5
Index: 4, Value: 6

========== Let Go infer the array length ==========
Colors: [Red Green Blue]

========== Arrays are initialized with zero values ==========
Zero Values: [0 0 0]

========== Arrays can be compared ==========
Arrays Equal: true

========== 2D Matrix ==========
[[1 2 3 4] [5 6 7 8] [9 10 11 12]]
1
7
12

========== Modify an Element ==========
[[1 2 3 4] [5 6 100 8] [9 10 11 12]]

========== Iterate Using nested loops ==========
matrix[0][0] = 1
matrix[0][1] = 2
matrix[0][2] = 3
matrix[0][3] = 4
matrix[1][0] = 5
matrix[1][1] = 6
matrix[1][2] = 100
matrix[1][3] = 8
matrix[2][0] = 9
matrix[2][1] = 10
matrix[2][2] = 11
matrix[2][3] = 12

========== Print row by row ==========
Matrix:
[1 2 3 4]
[5 6 100 8]
[9 10 11 12]

========== Show len() for a 2D array ==========
Rows: 3
Columns: 4

========== Show zero values for a 2D array ==========
Zero Matrix:
[0 0 0]
[0 0 0]

========== Sum of All Elements ==========
Sum: 171

========== Iterate Using Values ==========
1 2 3 4
5 6 100 8
9 10 11 12

========== Modify every element ==========
After Doubling:
[2 4 6 8]
[10 12 200 16]
[18 20 22 24]

========== Compare two matrices ==========
Matrices Equal: true

*/
