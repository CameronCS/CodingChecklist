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

var operations [5]byte

type Expression struct {
	left_operand  float64
	operation     byte
	right_operand float64
}

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

func AdvancedMainLoop() {
	operations = [5]byte{'+', '-', '*', '/', '^'}
	var err error
	var stdIn *bufio.Reader = bufio.NewReader(os.Stdin)

	fmt.Print("Enter an expression to have it evaluated: ")

	var expression string
	expression, err = stdIn.ReadString('\n')

	if err != nil {
		print("Error: ", err)
		return
	}

	expr, err := GetExpression(expression)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ans, err := expr.Evaluate()
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Printf("%v %s %v = %v", expr.left_operand, string(expr.operation), expr.right_operand, ans)
}

func IsSpace(char byte) bool {
	return char == ' '
}

func IsDot(char byte) bool {
	return char == '.'
}

func IsNumber(char byte, index int) bool {
	if IsDot(char) {
		return true
	}

	_, err := strconv.Atoi(string(char))
	return err == nil
}

func ConcatString(prev string, add string, dot bool, index int) (string, error) {
	if dot {
		return prev + ".", nil
	}
	if IsNumber(add[0], index) {
		return prev + add, nil
	}
	return "", errors.New(strings.Join([]string{"Invalid number", string(add), "at position", strconv.Itoa(index + 1)}, " "))
}

func GetExpression(expression string) (Expression, error) {
	var expr Expression
	var err error

	var has_dot bool = false
	var lhs string = ""
	var rhs string = ""
	var has_operation bool = false
	var isLHS bool = true
	for i := 0; i < len(expression); i++ {
		if expression[i] == '\r' || expression[i] == '\n' {
			continue
		}

		if IsSpace(expression[i]) {
			continue
		}

		var isNumber = IsNumber(expression[i], i)
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

			if !isOp {
				return expr, errors.New(strings.Join([]string{"invalid number", string(expression[i]), "at position", strconv.Itoa(i + 1)}, " "))
			}
		} else {
			if isLHS {
				var is_dot bool = IsDot(expression[i])
				if has_dot && is_dot {
					return expr, errors.New("too many '.'s in a single side of the operation (lhs)")
				}
				lhs, err = ConcatString(lhs, string(expression[i]), is_dot, i)
				if err != nil {
					return expr, err
				}
			} else {
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
	expr.left_operand, err = strconv.ParseFloat(lhs, 64)
	if err != nil {
		return expr, errors.New("issue With the left hand side of the expression")
	}

	expr.right_operand, err = strconv.ParseFloat(rhs, 64)
	if err != nil {
		return expr, errors.New("issue With the right hand side of the expression")
	}

	return expr, nil
}
