package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAlphabet(t *testing.T) {
	alphabetMap = createAlphabet(make(map[int]string))
	assert.Equal(t, alphabetMap[0], "A")
	assert.Equal(t, alphabetMap[25], "Z")
}

func TestCreateEncryptedAlphabet(t *testing.T) {
	alphabetMap = createAlphabet(make(map[int]string))
	alphabetMap = createEncryptedAlphabet(alphabetMap, 6)
	assert.Equal(t, alphabetMap[6], "A")
}

func TestCreateEncryptedAlphabetWithNegativeNumber(t *testing.T) {
	alphabetMap = createAlphabet(make(map[int]string))
	alphabetMap = createEncryptedAlphabet(alphabetMap, -1)
	assert.Equal(t, alphabetMap[0], "B")
}

func TestEncryptionUsingCesar(t *testing.T) {
	alphabetMap = createAlphabet(make(map[int]string))
	assert.Equal(t, EncryptUsingCesarMethod(alphabetMap, "CASA", 1), "DBTB")
	assert.Equal(t, EncryptUsingCesarMethod(alphabetMap, "ABECEDARIO", 2), "CDGEGFCTKQ")
}
