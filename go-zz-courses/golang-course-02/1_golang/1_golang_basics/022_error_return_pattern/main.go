package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// Go doesn't use exceptions for normal failures.
	// Functions return errors as normal return values.

	// value, err := someFunction()
	// if err != nil {
	//     handle the error
	// }

	if err := run(); err != nil {
		// log.Fatal prints the error and exits the program.
		log.Fatal(err)
	}
}

func run() error {
	input := "3"

	level, err := parseLevel(input)
	if err != nil {
		return err
	}

	fmt.Println("Selected Level:", level)
	return nil
}

func parseLevel(s string) (int, error) {
	// (value, error)
	// nil error     -> success
	// non-nil error -> failure

	n, err := strconv.Atoi(s)
	if err != nil {
		// %w wraps the original error
		return 0, fmt.Errorf("level must be a number: %w", err)
	}

	if n < 1 || n > 5 {
		return 0, fmt.Errorf("level must be between 1 and 5")
	}

	return n, nil
}
