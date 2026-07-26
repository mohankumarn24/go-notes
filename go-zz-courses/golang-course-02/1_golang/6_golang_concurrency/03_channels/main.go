package main

import (
	"fmt"
	"time"
)

func main() {

	// Channel:
	// A channel is a communication pipe between goroutines.
	// One goroutine sends data, another goroutine receives it.

	type User struct {
		ID   int
		Name string
	}

	// Create an unbuffered channel that can carry User values.
	ch := make(chan User)

	// Worker goroutine
	go func() {

		// Simulate a slow operation
		// (Database call, HTTP request, Microservice call, etc.)
		time.Sleep(200 * time.Millisecond)

		fmt.Println("Worker: sending user...")

		// Send a User into the channel.
		// Since this is an unbuffered channel,
		// the send waits until another goroutine receives it.
		ch <- User{
			ID:   100,
			Name: "Sangam",
		}

		fmt.Println("Worker: user sent")
	}()

	fmt.Println("Main: waiting to receive user...")

	// Receive a User from the channel.
	// If nothing has been sent yet, main blocks here.
	u := <-ch

	fmt.Println("Main: received user:", u)
	fmt.Println("ID:", u.ID)
	fmt.Println("Name:", u.Name)
}

/*
Main: waiting to receive user...
Worker: sending user...
Worker: user sent
Main: received user: {100 Sangam}
ID: 100
Name: Sangam
*/
