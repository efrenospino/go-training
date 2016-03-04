package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquesColorsFunction(t *testing.T) {
	name1, color1 := add("Efren")
	name2, color2 := add("Efren")
	_, color3 := add("Ana")
	assert.Equal(t, name1, name2, "Are the same names, Ok")
	assert.Equal(t, color1, color2, "Are the same names, because the colors too. Ok")
	assert.NotEqual(t, color1, color3, "Name diferents, colors diferents")
	assert.NotEqual(t, color2, color3, "Name diferents, colors diferents")
}
