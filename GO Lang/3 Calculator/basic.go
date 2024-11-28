package main

import (
	"fmt"
	"math"
	"os"
)

func BasicMainLoop() {
	var num1 float64
	fmt.Print("Enter a number: ")
	fmt.Scan(&num1)

	fmt.Print("Operations\n1) Add\n2) Subtract\n3) Multiply\n4) Divide\n5) Power\nSelect: ")
	var op int
	var charOp string
	fmt.Scan(&op)

	fmt.Print("Enter another number: ")
	var num2 float64
	fmt.Scan(&num2)

	var result float64
	switch op {
	case 1:
		result = num1 + num2
		charOp = "+"
	case 2:
		result = num1 - num2
		charOp = "-"
	case 3:
		result = num1 * num2
		charOp = "×"
	case 4:
		result = num1 / num2
		if num2 == 0 {
			fmt.Println("Can not divide by 0!")
			os.Exit(1)
		}
		charOp = "÷"
	case 5:
		charOp = "^"
		result = math.Pow(num1, num2)
	}

	fmt.Printf("%v %s %v = %v", num1, charOp, num2, result)
}
