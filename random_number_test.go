package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRandomNumber(t *testing.T) {
	numberProvider := MockRandom{}

	numberProvider.On("GetRandomNum").
		Once().
		Return(1, nil)

	result, err := numberProvider.GetRandomNum()

	numberProvider.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, 1, result)
}
