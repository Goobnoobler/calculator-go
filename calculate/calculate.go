package calculate

import (
	"errors"
	"strconv"
	"strings"
)

//Sentinal errors

var ErrUnmatchedParen = errors.New("unmatched closing parenthesis detected")
var ErrUnclosedParen = errors.New("unclosed parenthesis detected")
var ErrUnexpectedChar = errors.New("calculator only accepts digits and +-*/")

var ErrDivideByZero = errors.New("cannot divide by zero")
var ErrUnexpectedOperator = errors.New("unknown operator in expression")
var ErrNotEnoughOperands = errors.New("not enough operands for operater")
var ErrTooManyValues = errors.New("expression left more than one value")
var ErrNoValue = errors.New("no answer returned")

// The aim of this function is to convert human-readable 2*(4+3) format to reverse Polish notation 2,4,3,+,*
// It follows the shunting yard algorithm which is used regularly for software-based calculators
func Shunt(expression string) ([]string, error) {
	var output []string
	var operators []byte
	var token strings.Builder
	var prec = map[byte]int{'+': 1, '-': 1, '*': 2, '/': 2}
	var err error

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
		} else if expression[i] == '(' {
			if token.Len() > 0 {
				output = append(output, token.String())
				token.Reset()
			}
			operators = append(operators, expression[i])
		} else if expression[i] == ')' {
			if token.Len() > 0 {
				output = append(output, token.String())
				token.Reset()
			}
			for len(operators) > 0 && operators[len(operators)-1] != '(' {
				output = append(output, string(operators[len(operators)-1]))
				operators = operators[:len(operators)-1]
			}
			if len(operators) == 0 {
				err = ErrUnmatchedParen
				return nil, err
			}
			operators = operators[:len(operators)-1]
		} else {
			err = ErrUnexpectedChar
			return nil, err
		}
	}
	if token.Len() > 0 {
		output = append(output, token.String())
		token.Reset()
	}

	for i := len(operators); i > 0; i-- {
		if operators[i-1] == '(' {
			err = ErrUnclosedParen
			return nil, err
		}
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
					if num2 == 0 {
						err = ErrDivideByZero
						return 0, err
					}
					calc = calc[:len(calc)-2]
					calc = append(calc, num1/num2)
				default:
					err = ErrUnexpectedOperator
					return 0, err
				}
			} else {
				err = ErrNotEnoughOperands
				return 0, err
			}

		} else {
			calc = append(calc, ans)
		}
	}
	if len(calc) == 1 {
		return calc[0], nil
	} else if len(calc) == 0 {
		err = ErrNoValue
		return 0, err
	} else {
		err = ErrTooManyValues
		return 0, err
	}

}
