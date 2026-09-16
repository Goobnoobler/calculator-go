package calculate

import (
	"errors"
	"slices"
	"testing"
)

func TestShunt(t *testing.T) {

	tests := []struct {
		name    string
		input   string
		want    []string
		wantErr error
	}{
		{"basic addition", "2+3", []string{"2", "3", "+"}, nil},
		{"parenthesis", "2*(4+3)", []string{"2", "4", "3", "+", "*"}, nil},
		{"popping", "2*4+3", []string{"2", "4", "*", "3", "+"}, nil},
		{"multi digit", "12+345", []string{"12", "345", "+"}, nil},
		{"decimal", "12+3.45", []string{"12", "3.45", "+"}, nil},
		{"umatched parenthesis", "2*4+3)", nil, ErrUnmatchedParen},
		{"unclosed parenthesis", "2*(4+3", nil, ErrUnclosedParen},
		{"unexpected character", "2*4+3$", nil, ErrUnexpectedChar},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Shunt(tt.input)

			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("Shunt(%v) = %v, want %v", tt.input, err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Fatalf("Shunt(%v) returned unexpected error: %v", tt.input, err)
			}

			if !slices.Equal(got, tt.want) {
				t.Errorf("Shunt(%v) = %v, want %v", tt.input, got, tt.want)
			}

		})
	}

}

func TestEval(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    float64
		wantErr error
	}{
		{"addition", []string{"2", "3", "+"}, 5, nil},
		{"subtraction", []string{"9", "4", "-"}, 5, nil},
		{"multiplication", []string{"4", "5", "*"}, 20, nil},
		{"division", []string{"20", "2", "/"}, 10, nil},
		{"divide by zero", []string{"20", "0", "/"}, 0, ErrDivideByZero},
		{"unexpected operator", []string{"20", "5", "%"}, 0, ErrUnexpectedOperator},
		{"not enough operands", []string{"20", "+"}, 0, ErrNotEnoughOperands},
		{"too many values", []string{"20", "5"}, 0, ErrTooManyValues},
		{"no value", []string{}, 0, ErrNoValue},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Eval(tt.input)

			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("Eval(%v) = %v, want %v", tt.input, err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Fatalf("Eval(%v) returned unexpected error: %v", tt.input, err)
			}

			if got != tt.want {
				t.Errorf("Eval(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
