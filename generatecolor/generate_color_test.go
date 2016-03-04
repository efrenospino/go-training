package generatecolor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateColorByName(t *testing.T) {
	assert.Equal(t, generateValuesByLettersInName("efren"), "#c145c8") // "935929243"
}
