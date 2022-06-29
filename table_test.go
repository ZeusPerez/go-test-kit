package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumTable(t *testing.T) {
	type test struct {
		name     string
		input    [2]int
		expected int
	}

	tests := []test{
		{name: "Happy Path", input: [2]int{10, 25}, expected: 35},
		{name: "Negative", input: [2]int{10, 25}, expected: 35},
		{name: "Mixed", input: [2]int{10, 25}, expected: 35},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := sum(tc.input[0], tc.input[1])
			assert.Equal(t, tc.expected, result)
		})
	}
}
