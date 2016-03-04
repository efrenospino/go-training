package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAreIsomorphicsFunction(t *testing.T) {
	assert.Equal(t, areIsomorphics("foo", "bar"), false)
	assert.Equal(t, areIsomorphics("paper", "title"), true)
	assert.Equal(t, areIsomorphics("egg", "add"), true)
}
