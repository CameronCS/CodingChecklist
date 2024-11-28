package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Println("Invalid args supplied")
		return
	}

	if args[1] == "basic" {
		BasicMainLoop()
	} else if args[1] == "adv" {
		AdvancedMainLoop()
	}
}
