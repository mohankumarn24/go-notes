package main

import "fmt"

func main() {
	// map[keyType]valueType
	ages := map[string]int{
		"Armani": 65,
		"Ford":   35,
	}

	fmt.Println("Armani's Age:", ages["Armani"]) // Armani's Age: 65
	fmt.Println("Length:", len(ages))            // Length: 2

	// Declare a nil map (read is safe, write causes panic)
	var scores map[string]int

	fmt.Println("Nil Map:", scores)                        // Nil Map: map[]
	fmt.Println("Missing Key:", scores["Math"])            // Missing Key: 0 -> Returns zero value
	score1, exists1 := scores["Math"]                      // point of confusion. use this instead
	fmt.Printf("Score: %d, Exists: %t\n", score1, exists1) // Score: 0, Exists: false
	// value, ok :=  scores["Math"]
	// _, ok := scores["Math"]

	// Create a map using make
	scores = make(map[string]int)

	scores["Math"] = 90
	scores["Science"] = 95

	fmt.Println("Scores:", scores) // Scores: map[Math:90 Science:95]

	// Update a value
	scores["Math"] = 99
	fmt.Println("Updated Scores:", scores) // Updated Scores: map[Math:99 Science:95]

	// Check if a key exists (comma-ok idiom)
	score, exists := scores["Math"]
	fmt.Println("Math:", score, "Exists:", exists) // Math: 99 Exists: true

	score, exists = scores["English"]
	fmt.Println("English:", score, "Exists:", exists) // English: 0 Exists: false

	// Iterate over a map
	for subject, score := range scores {
		fmt.Println(subject, score)
	}
	// Math 99
	// Science 95

	// Iterate over only keys
	for subject := range scores {
		fmt.Println("Subject:", subject)
	}
	// Subject: Math
	// Subject: Science

	// Iterate over only values
	for _, score := range scores {
		fmt.Println("Score:", score)
	}
	// Score: 99
	// Score: 95

	users := map[string]string{
		"u1": "Armani",
		"u2": "John",
		"u3": "Rahul",
	}

	fmt.Println("Users:", users) // Users: map[u1:Armani u2:John u3:Rahul]

	// Delete entries
	delete(users, "u2")
	delete(users, "u100") // No error if the key doesn't exist

	fmt.Println("Users After Delete:", users) // Users After Delete: map[u1:Armani u3:Rahul]

	// Clear all entries (Go 1.21+)
	clear(users)
	fmt.Println("Users After Clear:", users) // Users After Clear: map[]
}
