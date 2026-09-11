package main

import (
	"errors"
	"fmt"
	"unicode"
)

// "log"
// "net/http"

// "github.com/goobnoobler/server/server"
//

func main() {
	expression := []rune("3+4*2-1")
	EvalExpression(expression)
	output, err := EvalExpression(expression)
	fmt.Printf("%s %v\n", string(output), err)
}

func EvalExpression(expression []rune) ([]rune, error) {
	var output []rune
	var operators []rune

	for i := range expression {
		if unicode.IsDigit(expression[i]) {
			output = append(output, expression[i])
		} else if expression[i] == '+' || expression[i] == '-' || expression[i] == '*' || expression[i] == '/' {
			var prec = map[rune]int{'+': 1, '-': 1, '*': 2, '/': 2}
			for len(operators) > 0 && prec[operators[len(operators)-1]] >= prec[expression[i]] {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1] // shrinking it by 1
			}
			operators = append(operators, expression[i])
		} else {
			return nil, errors.New("Calculator only accepts Digits and +-*/.")
		}
	}
	for i := len(operators); i > 0; i-- {
		output = append(output, operators[i-1])
	}
	return output, nil
}
