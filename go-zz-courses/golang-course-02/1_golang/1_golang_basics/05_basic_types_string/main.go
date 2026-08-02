package main

import (
	"fmt"
	"strings"
)

func main() {
	// =============================================
	// ===== 1. String Mutability/Immutability =====
	// =============================================

	// === Strings are immutable ===
	str1 := "abcd"
	str2 := str1
	bytes := []byte(str1) // Convert the string to a mutable byte slice. []byte(str1) creates a copy of the string's bytes
	bytes[0] = 'A'
	// str1 = "Abcd"
	str1 = string(bytes)           // Convert the modified byte slice back to a new string. string(bytes) creates a new string
	fmt.Println("String 1:", str1) // Abcd
	fmt.Println("String 2:", str2) // abcd (unchanged)

	// === Arrays are mutable (value types) ===
	names1 := [3]string{"Sachin", "Virat", "Dhoni"} // Arrays are mutable value types
	names2 := names1                                // Assignment creates a copy of the entire array
	names1[0] = "Rohit"                             // Modify the original array
	fmt.Println("names1:", names1)                  // [Rohit Virat Dhoni]
	fmt.Println("names2:", names2)                  // [Sachin Virat Dhoni]

	// === Slices are mutable (shared underlying array) ===
	slice1 := []string{"Sachin", "Virat", "Dhoni"}
	slice2 := slice1 // Assignment shares the same underlying array
	slice1[0] = "Rohit"
	fmt.Println("slice1:", slice1) // [Rohit Virat Dhoni]
	fmt.Println("slice2:", slice2) // [Rohit Virat Dhoni]

	// === Slice (Independent Copy using make + copy) ===
	newNames1 := []string{"Sachin", "Virat", "Dhoni"} // Create a slice
	newNames2 := make([]string, len(newNames1))       // Create another slice with the same length
	copy(newNames2, newNames1)                        // Copy the elements
	newNames1[0] = "Rohit"                            // Modify the original slice
	fmt.Println("newNames1:", newNames1)              // [Rohit Virat Dhoni]
	fmt.Println("newNames2:", newNames2)              // [Sachin Virat Dhoni]

	// =============================================
	// ========= 2. Common String functions ========
	// =============================================
	firstName := "Emporio"
	lastName := "Armani"

	// String concatenation
	fullName := firstName + " " + lastName
	fmt.Println("Full Name:", fullName) // Emporio Armani

	// Common string functions [start:end]
	fmt.Println("Substring [0:6]:", firstName[0:6]) // Empori
	fmt.Println("Substring [3:]:", firstName[3:])   // orio
	fmt.Println("Substring [2:5]:", firstName[2:5]) // por

	fmt.Println("HasPrefix():", strings.HasPrefix(fullName, "Emp"))   // true
	fmt.Println("HasSuffix():", strings.HasSuffix(fullName, "ni"))    // true
	fmt.Println("Contains():", strings.Contains(fullName, "Emporio")) // true

	fmt.Println("Index(\"m\"):", strings.Index(fullName, "m"))           // 1
	fmt.Println("Index(\"rio\"):", strings.Index(fullName, "rio"))       // 4
	fmt.Println("Index(\"Sachin\"):", strings.Index(fullName, "Sachin")) // -1
	fmt.Println("LastIndex(\"a\"):", strings.LastIndex(fullName, "a"))   //11

	afterSplit := strings.Split(fullName, " ")
	afterJoin := strings.Join(afterSplit, " ")
	fmt.Println("Split():", afterSplit) //[Emporio Armani]
	fmt.Println("Join():", afterJoin)   // Emporio Armani

	fmt.Println("Count(\"a\"):", strings.Count(fullName, "a"))               // 1
	fmt.Println("ToUpper():", strings.ToUpper(fullName))                     // EMPORIO ARMANI
	fmt.Println("ToLower():", strings.ToLower(fullName))                     // emporio armani
	fmt.Println("Compare(\"abc\", \"xyz\"):", strings.Compare("abc", "xyz")) // -1
	fmt.Println("Fields():", strings.Fields("  a  b   c "))                  // [a b c]
	fmt.Println("Repeat():", strings.Repeat("*", 10))                        // **********

	fmt.Println("Replace():", strings.Replace(fullName, "a", "@", 1))                // Emporio Arm@ni
	fmt.Println("ReplaceAll():", strings.ReplaceAll(fullName, "Emporio", "Giorgio")) // Giorgio Armani

	fmt.Println("TrimSpace():", strings.TrimSpace("  Hello  "))          // Hello
	fmt.Println("TrimPrefix():", strings.TrimPrefix(fullName, "Emp"))    // orio Armani
	fmt.Println("TrimSuffix():", strings.TrimSuffix(fullName, "Armani")) // Emporio

	// String comparison
	fmt.Println("Equals (==):", fullName == "Emporio Armani") // true
	fmt.Println("Equals (==):", fullName == "Giorgio Armani") // false

	// Case-insensitive string comparison
	fmt.Println("EqualFold():", strings.EqualFold("GO", "go"))           // true
	fmt.Println("EqualFold():", strings.EqualFold("Emporio", "EMPORIO")) // true
	fmt.Println("EqualFold():", strings.EqualFold("Emporio", "Armani"))  // false

	// =============================================
	// ========== 3. String Concatenation ==========
	// =============================================
	// 1. strings.Builder (efficient for many concatenations, especially in loops)
	var builder strings.Builder
	builder.WriteString("Emporio")
	builder.WriteString(" ")
	builder.WriteString("Armani")
	fmt.Println("string concat using strings.Builder:", builder.String()) // Emporio Armani

	// 2. + operator (most common for a few strings)
	firstName2 := "Emporio"
	lastName2 := "Armani"
	fullName2 := firstName2 + " " + lastName2
	fmt.Println("string concat using +:", fullName2) // Emporio Armani

	// 3. += operator (append to an existing string)
	name3 := "Emporio"
	name3 += " Armani"
	fmt.Println("string concat using +=:", name3) // Emporio Armani

	// 4. fmt.Sprintf() (format strings with variables)
	firstName4 := "Emporio"
	lastName4 := "Armani"
	age4 := 21
	result4 := fmt.Sprintf("%s %s is %d years old", firstName4, lastName4, age4)
	fmt.Println("string concat using fmt.Sprintf():", result4) // Emporio Armani is 21 years old

	// 5. strings.Join() (join a slice of strings with a separator)
	names5 := []string{"Emporio", "Armani"}
	fullName5 := strings.Join(names5, " ")
	fmt.Println("string concat using strings.Join():", fullName5) // Emporio Armani
}
