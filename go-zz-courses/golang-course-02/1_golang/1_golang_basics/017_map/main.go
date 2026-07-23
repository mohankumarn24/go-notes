package main

import "fmt"

func main() {
	// map[keyType]valueType
	ages := map[string]int{
		"Sangam": 65,
		"John":   35,
	}

	fmt.Println("Sangam's Age:", ages["Sangam"])
	fmt.Println("Length:", len(ages))

	// Declare a nil map (read is safe, write causes panic)
	var scores map[string]int

	fmt.Println("Nil Map:", scores)
	fmt.Println("Missing Key:", scores["Math"]) // Returns zero value

	// Create a map using make
	scores = make(map[string]int)

	scores["Math"] = 90
	scores["Science"] = 95

	fmt.Println("Scores:", scores)

	// Update a value
	scores["Math"] = 99
	fmt.Println("Updated Scores:", scores)

	// Check if a key exists (comma-ok idiom)
	score, exists := scores["Math"]
	fmt.Println("Math:", score, "Exists:", exists)

	score, exists = scores["English"]
	fmt.Println("English:", score, "Exists:", exists)

	// Iterate over a map
	for subject, score := range scores {
		fmt.Println(subject, score)
	}

	// Iterate over only keys
	for subject := range scores {
		fmt.Println("Subject:", subject)
	}

	// Iterate over only values
	for _, score := range scores {
		fmt.Println("Score:", score)
	}

	users := map[string]string{
		"u1": "Sangam",
		"u2": "John",
		"u3": "Rahul",
	}

	fmt.Println("Users:", users)

	// Delete entries
	delete(users, "u2")
	delete(users, "u100") // No error if the key doesn't exist

	fmt.Println("Users After Delete:", users)

	// Clear all entries (Go 1.21+)
	clear(users)
	fmt.Println("Users After Clear:", users)
}
