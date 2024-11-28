package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Create the operation array
var operations [5]byte

// create the struct for our expression
type Expression struct {
	left_operand  float64
	operation     byte
	right_operand float64
}

// Simple evaluate method for the struct
func (e Expression) Evaluate() (float64, error) {
	switch e.operation {
	case '+':
		return e.left_operand + e.right_operand, nil
	case '-':
		return e.left_operand - e.right_operand, nil
	case '*':
		return e.left_operand * e.right_operand, nil
	case '/':
		if e.right_operand == 0 {
			return 0, errors.New("zero division error")
		}
		return e.left_operand / e.right_operand, nil

	case '^':
		return math.Pow(e.left_operand, e.right_operand), nil
	}
	return 0, errors.New("invalid operation expression")
}

// Loop for the advanced calculator
func AdvancedMainLoop() {
	// Initialise the byte array
	operations = [5]byte{'+', '-', '*', '/', '^'}
	// Define an error for the loop
	var err error
	// Create the stdIn for the reader
	var stdIn *bufio.Reader = bufio.NewReader(os.Stdin)
	// Ask the user for their expression
	fmt.Print("Enter an expression to have it evaluated: ")
	// Define the expression string
	var expression string
	// Read the expression in
	expression, err = stdIn.ReadString('\n')
	// Check for an error
	if err != nil {
		print("Error: ", err)
		return
	}
	// Get the expression struct from the expression string
	expr, err := GetExpression(expression)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Get the answer by evaluating the expression
	ans, err := expr.Evaluate()
	// Check for error
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	// print the expression
	fmt.Printf("%v %s %v = %v", expr.left_operand, string(expr.operation), expr.right_operand, ans)
}

// Func to check if the current char is a space
func IsSpace(char byte) bool {
	return char == ' '
}

// Check to see if the current char is a dot
func IsDot(char byte) bool {
	return char == '.'
}

// Check to see if the current char is a number
func IsNumber(char byte, index int) bool {
	// If its a dot it technically a part of a number
	if IsDot(char) {
		return true
	}
	// check to see if the conversion throws an error
	_, err := strconv.Atoi(string(char))
	// return true or false depending on if there is an error
	return err == nil
}

// This method is a simple method to concatonate a string
func ConcatString(prev string, add string, dot bool, index int) (string, error) {
	if dot {
		return prev + ".", nil
	}
	if IsNumber(add[0], index) {
		return prev + add, nil
	}
	return "", errors.New(strings.Join([]string{"Invalid number", string(add), "at position", strconv.Itoa(index + 1)}, " "))
}

// This method gets the expression string from the expression string
func GetExpression(expression string) (Expression, error) {
	// Define the variables
	var expr Expression
	var err error

	var has_dot bool = false
	var lhs string = ""
	var rhs string = ""
	var has_operation bool = false
	var isLHS bool = true

	// Loop through each charactor of the application
	for i := 0; i < len(expression); i++ {
		if expression[i] == '\r' || expression[i] == '\n' {
			continue
		}

		// ignore space
		if IsSpace(expression[i]) {
			continue
		}

		// Determine if its a number
		var isNumber = IsNumber(expression[i], i)
		// If its not we assume its an opperand
		if !isNumber {
			var isOp = false
			for j := 0; j < len(operations); j++ {
				if expression[i] == operations[j] {
					if has_operation {
						return expr, errors.New("too many operations supplied for this expression")
					}
					isOp = true
					has_operation = true
					expr.operation = expression[i]
					isLHS = false
					has_dot = false
				}
			}
			// If its not an operannd we return an error
			if !isOp {
				return expr, errors.New(strings.Join([]string{"invalid number", string(expression[i]), "at position", strconv.Itoa(i + 1)}, " "))
			}
		} else {
			// if it is a number we come here
			// we check if we are busy with the left hand side
			if isLHS {
				// check if its a dot
				var is_dot bool = IsDot(expression[i])
				// check if we have a dot and it is a dot
				if has_dot && is_dot {
					// if it is a dot and we have a dot we return an error
					return expr, errors.New("too many '.'s in a single side of the operation (lhs)")
				}
				// otherwise we concatonate the string with the expression term
				lhs, err = ConcatString(lhs, string(expression[i]), is_dot, i)
				// we check if that gives us an error and if it does we return it
				if err != nil {
					return expr, err
				}
			} else {
				// the same is done here except with the righthand side
				var is_dot bool = IsDot(expression[i])
				if has_dot && is_dot {
					return expr, errors.New("too many '.'s in a single side of the operation (lhs)")
				}
				rhs, err = ConcatString(rhs, string(expression[i]), is_dot, i)
				if err != nil {
					return expr, err
				}
			}
		}
	}
	// we assign the left opperand and check for error otherwise return the error
	expr.left_operand, err = strconv.ParseFloat(lhs, 64)
	if err != nil {
		return expr, errors.New("issue With the left hand side of the expression")
	}

	// we assign the right opperand and check for error otherwise return the error
	expr.right_operand, err = strconv.ParseFloat(rhs, 64)
	if err != nil {
		return expr, errors.New("issue With the right hand side of the expression")
	}

	// we return the expression
	return expr, nil
}
