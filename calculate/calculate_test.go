package calculate

import (
	"errors"
	"testing"
)

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
		// your turn: multiplication and division
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
