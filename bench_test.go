package main

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano())
		x := rand.Int()
		y := rand.Int()
		sum(x, y)
	}
}
