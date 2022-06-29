package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	result := sum(10, 25)
	expected := 35

	assert.Equal(t, expected, result)
}

func TestSumNegative(t *testing.T) {
	result := sum(-5, -7)
	expected := -12

	assert.Equal(t, expected, result)
}

func TestSumMixed(t *testing.T) {
	result := sum(-14, 28)
	expected := 14

	assert.Equal(t, expected, result)
}
