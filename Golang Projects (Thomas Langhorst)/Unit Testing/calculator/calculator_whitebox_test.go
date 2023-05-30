package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var isZeroTestcases = []struct {
	name string

	expected bool
	arg      float64
}{
	{"argument is zero", true, 0.0},
	{"argument is not zero", false, 1.0},
}

func TestIsZero(t *testing.T) {
	for _, tc := range isZeroTestcases {
		t.Run(tc.name, func(t *testing.T) {

			assert := assert.New(t)
			got := isZero(tc.arg)

			assert.Equal(tc.expected, got)
		})
	}
}
