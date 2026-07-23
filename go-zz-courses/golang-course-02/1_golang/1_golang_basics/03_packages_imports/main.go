package main

// import brings packages into the current file.
// Imported packages are available only in the file where they are imported.

import (
	"fmt"
	"math"
)

func main() {
	// packageName.FunctionName() calls an exported function from a package.
	fmt.Println("Sqrt(25):", math.Sqrt(25))
}
