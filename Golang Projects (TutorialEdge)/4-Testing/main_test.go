package main

import (
	"fmt"
	"testing"
)

func Test_Calculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}

func Test_Table_Calculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{9999, 10001},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			fmt.Printf("Test Failed: {%v} inputted / {%v} tested / {%v} received\n", test.input, test.expected, output)
		}
	}
}
