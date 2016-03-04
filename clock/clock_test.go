package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClockFunction(t *testing.T) {
	assert.Equal(t, calculateAngle(15, 8), 46.0)
	assert.Equal(t, calculateAngle(9, 0), 270.0)
	assert.Equal(t, calculateAngle(-9, 0), 270.0)
	assert.Equal(t, calculateAngle(-15, 8), 46.0)
	assert.Equal(t, calculateAngle(15, -8), 46.0)
	assert.Equal(t, calculateAngle(-15, -8), 46.0)
	assert.Equal(t, calculateAngle(-9, -0), 270.0)
}
