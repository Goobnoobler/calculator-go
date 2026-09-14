package calculate

import (
	"errors"
	"strconv"
	"strings"
)

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
				err = errors.New("unmatched closing parenthesis detected")
				return nil, err
			}
			operators = operators[:len(operators)-1]
		} else {
			err = errors.New("calculator only accepts digits and +-*/")
			return nil, err
		}
	}
	if token.Len() > 0 {
		output = append(output, token.String())
		token.Reset()
	}

	for i := len(operators); i > 0; i-- {
		if operators[i-1] == '(' {
			err = errors.New("unclosed parenthesis detected")
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
						err = errors.New("cannot divide by zero")
						return 0, err
					}
					calc = calc[:len(calc)-2]
					calc = append(calc, num1/num2)
				default:
					err = errors.New("unknown character in expression")
					return 0, err
				}
			} else {
				err = errors.New("expression is too short")
				return 0, err
			}

		} else {
			calc = append(calc, ans)
		}
	}
	if len(calc) == 1 {
		return calc[0], nil
	} else {
		err = errors.New("answer returned 2 values")
		return 0, err
	}

}
