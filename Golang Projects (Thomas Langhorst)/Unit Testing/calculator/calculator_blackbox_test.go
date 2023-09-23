package calculator_test

import (
	"errors"
	"testing"
	"unittesting/calculator"

	"github.com/stretchr/testify/assert"
)

var testcases = []struct {
	name string

	expected      float64
	expectedError error

	divisor float64
}{
	{"division", 2.0, nil, 5.0},
	{"division by negative value", -2.0, nil, -5.0},
	{"division by zero", 0.0, errors.New("division by zero"), 0.0},
}

func TestDivide(t *testing.T) {
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			gotValue, gotError := calculator.Divide(10.0, tc.divisor)

			if gotError != nil {
				if gotError.Error() != tc.expectedError.Error() {
					t.Errorf("expected error: %s, got error: %s", tc.expectedError.Error(), gotError.Error())
				}
			}

			if gotValue != tc.expected {
				t.Errorf("expected: %.1f, got: %1.f", tc.expected, gotValue)
			}
		})
	}
}

func TestDivideTestify(t *testing.T) {
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			assert := assert.New(t)

			gotValue, gotError := calculator.Divide(10.0, tc.divisor)

			assert.Equal(tc.expected, gotValue)
			assert.Equal(tc.expectedError, gotError)
			assert.NotEqual("division by zero", gotError)

			if gotError != nil {
				if gotError.Error() != tc.expectedError.Error() {
					t.Errorf("expected error: %s, got error: %s", tc.expectedError.Error(), gotError.Error())
				}
			}

			if gotValue != tc.expected {
				t.Errorf("expected: %.1f, got: %1.f", tc.expected, gotValue)
			}
		})
	}
}
