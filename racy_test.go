package main

import (
	"context"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRacy(t *testing.T) {
	mockRand := &MockRandom{}
	mockRand.On("GetRandomNum").Times(200).Return(10, nil)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			_, err := randomSum(context.TODO(), mockRand)
			assert.NoError(t, err)
		}()
	}

}
