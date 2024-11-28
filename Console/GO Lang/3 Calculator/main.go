package main

import (
	"fmt"
	"os"
)

func main() {
	// Get command line arguments
	args := os.Args

	// Check if length of args is not 2
	// If its not we are unable to run the application
	if len(args) != 2 {
		fmt.Println("Invalid args supplied")
		return
	}

	// We check to see if the only arg supplied is basic or adv and select the appropriate calculator
	if args[1] == "basic" {
		BasicMainLoop()
	} else if args[1] == "adv" {
		AdvancedMainLoop()
	}
}
