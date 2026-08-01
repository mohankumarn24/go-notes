package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Outer Loop (i):", i)

		for j := 0; j < 10; j++ {
			if j == 5 {
				fmt.Println("  Break at j =", j)
				break
				// continue
			}

			fmt.Printf("  i=%d, j=%d\n", i, j)
		}

		fmt.Println()
	}
}
