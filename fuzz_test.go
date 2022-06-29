package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func FuzzSum(f *testing.F) {
	f.Add(0, int(time.Now().UnixNano())) // Seed corpus
	f.Fuzz(func(t *testing.T, x int, y int) {
		rand.Seed(time.Now().UnixNano())
		result := sum(x, y)
		assert.Equal(t, x+y, result)
	})
}
