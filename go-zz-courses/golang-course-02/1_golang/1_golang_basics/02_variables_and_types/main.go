package main

import "fmt"

func main() {
	// var variableName type = value
	var year int = 2026
	fmt.Println("Year:", year)

	var rating float64 = 4.8
	fmt.Println("Rating:", rating)

	// Declaration without initialization
	// Zero value is assigned automatically.
	var channelName string
	channelName = "AceDevHub"
	fmt.Println("Channel:", channelName)
}
