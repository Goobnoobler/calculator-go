package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// "log"
// "net/http"

// "github.com/goobnoobler/server/server"
//

func main() {
	var input = "0-100"

	output, err := Shunt(input)

	output2, err2 := Eval(output)
	if err == nil && err2 == nil {
		fmt.Printf("%v \n", output)
		fmt.Printf("%v \n", output2)
	} else {
		fmt.Printf("%v \n", err)
		fmt.Printf("%v \n", err2)
	}

}

func Shunt(expression string) ([]string, error) {
	var output []string
	var operators []byte
	var token strings.Builder
	var prec = map[byte]int{'+': 1, '-': 1, '*': 2, '/': 2}

	for i := range expression {
		if (expression[i] >= '0' && expression[i] <= '9') || expression[i] == '.' {
			token.WriteByte(expression[i])
		} else if expression[i] == '+' || expression[i] == '-' || expression[i] == '*' || expression[i] == '/' {
			if token.Len() > 0 {
				output = append(output, token.String())
				token.Reset()
			}
			for len(operators) > 0 && prec[operators[len(operators)-1]] >= prec[expression[i]] {
				output = append(output, string(operators[len(operators)-1]))
				operators = operators[:len(operators)-1] // shrinking it by 1
			}
			operators = append(operators, expression[i])
		} else {
			fmt.Println("Calculator only accepts Digits and +-*/.")
		}
	}
	if token.Len() > 0 {
		output = append(output, token.String())
		token.Reset()
	}

	for i := len(operators); i > 0; i-- {
		output = append(output, string(operators[i-1]))
	}
	return output, nil
}

func Eval(r []string) (float64, error) {

	var calc []float64
	var err error

	for i := range r {
		ans, parseErr := strconv.ParseFloat(r[i], 64)

		if parseErr != nil {
			if len(calc) > 1 {
				var num1 = calc[len(calc)-2]
				var num2 = calc[len(calc)-1]
				var op string = r[i]

				switch op {
				case "+":
					calc = calc[:len(calc)-2]
					calc = append(calc, num1+num2)

				case "-":
					calc = calc[:len(calc)-2]
					calc = append(calc, num1-num2)

				case "*":
					calc = calc[:len(calc)-2]
					calc = append(calc, num1*num2)

				case "/":
					calc = calc[:len(calc)-2]
					check0 := num1 / num2
					if math.IsInf(check0, 0) {
						err = errors.New("Cannot Divide by Zero.")
						return 0, err
					}
					calc = append(calc, num1/num2)
				default:
					err = errors.New("Unknown Character in Expression.")
					return 0, err
				}
			} else {
				err = errors.New("expression is too short.")
				return 0, err
			}

		} else {
			calc = append(calc, ans)
		}
	}
	if len(calc) == 1 {
		return calc[0], nil
	} else {
		err = errors.New("Answer returned 2 values.")
		return 0, err
	}

}
