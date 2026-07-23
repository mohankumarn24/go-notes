package main

import (
	"fmt"
	"strings"
)

func main() {
	firstName := "Sangam"
	lastName := "Mukherjee"

	// String concatenation
	fullName := firstName + " " + lastName

	fmt.Println("Full Name:", fullName)

	// Common string functions
	fmt.Println("ToUpper():", strings.ToUpper(fullName))
	fmt.Println("ToLower():", strings.ToLower(fullName))
	fmt.Println("Contains():", strings.Contains(fullName, "Sangam"))
	fmt.Println("HasPrefix():", strings.HasPrefix(fullName, "San"))
	fmt.Println("HasSuffix():", strings.HasSuffix(fullName, "jee"))
	fmt.Println("ReplaceAll():", strings.ReplaceAll(fullName, "Sangam", "John"))
	fmt.Println("Split():", strings.Split(fullName, " "))
	fmt.Println("TrimSpace():", strings.TrimSpace("  Hello  "))
	fmt.Println("Repeat():", strings.Repeat("*", 10))
}
