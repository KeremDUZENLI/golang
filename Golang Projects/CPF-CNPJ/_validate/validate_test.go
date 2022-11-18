package validate

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CpfOrCnpj(t *testing.T) {
	var tests = []struct {
		insert   string
		expected int
	}{
		{"79885060057", 1},
		{"241.199.910-08", 1},
		{"111.111.111-11", 0},

		{"44997315000120", 2},
		{"24.262.071/0001-20", 2},
		{"11.111.111/1111-11", 0},
	}

	for _, test := range tests {
		if output := CpfOrCnpj(test.insert); output != test.expected {
			fmt.Println("Test Failed")
		}
	}
}

func Test_Unit_CpfOrCnpj(t *testing.T) {
	var validationTestCases = []struct {
		input  string
		output int
	}{
		{"46589420890", 1},
		{"241.199.910-08", 1},
		{"241.199.B1A-08", 0},

		{"44997315000120", 2},
		{"89.747.580/0001-99", 2},
		{"36.382.680/0001-99", 0},
	}

	t.Run("Should return success", func(t *testing.T) {
		for _, validationTest := range validationTestCases {
			testErrorMessage := fmt.Sprintf("Testing: %s Expected: %v", validationTest.input, validationTest.output)

			checkResult := CpfOrCnpj(validationTest.input)
			assert.Equal(t, validationTest.output, checkResult, testErrorMessage)
		}
	})
}
