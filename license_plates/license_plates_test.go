package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountRangeBetweenNumbers(t *testing.T) {
	assert.Equal(t, countRangeBetweenNumbers(5, 746), 741)
	assert.Equal(t, countRangeBetweenNumbers(564, 353), 211)
}

func TestGetIndexOfLetter(t *testing.T) {
	alphabetMap = createAlphabet(make(map[int]string))
	assert.Equal(t, getIndexOfLetter("F", alphabetMap), 5)
	assert.Equal(t, getIndexOfLetter("Z", alphabetMap), 25)
}

func TestCountRangeBetweenLetters(t *testing.T) {
	assert.Equal(t, countRangeBetweenLetters("AAX", "AAA"), 23)
	assert.Equal(t, countRangeBetweenLetters("AAA", "ABA"), 26)
	assert.Equal(t, countRangeBetweenLetters("ACA", "AAC"), 54)
	assert.Equal(t, countRangeBetweenLetters("ACA", "ACC"), 2)
	assert.Equal(t, countRangeBetweenLetters("AZA", "AAA"), 650)
	assert.Equal(t, countRangeBetweenLetters("AZZ", "AAA"), 675)
	assert.Equal(t, countRangeBetweenLetters("BAA", "AAA"), 676)
	assert.Equal(t, countRangeBetweenLetters("BBB", "AAA"), 703)
}

func TestCountRangeBetweenLicensePlates(t *testing.T) {
	assert.Equal(t, countRangeBetweenLicencePlates("AAA-000", "AAD-000"), 3000)
	assert.Equal(t, countRangeBetweenLicencePlates("AAA-000", "AAD-999"), 3999)
	assert.Equal(t, countRangeBetweenLicencePlates("AAA-999", "AAD-000"), 2001)
	assert.Equal(t, countRangeBetweenLicencePlates("AAD-999", "AAA-000"), 3999)
	assert.Equal(t, countRangeBetweenLicencePlates("AAA-000", "ABA-000"), 26000)
	assert.Equal(t, countRangeBetweenLicencePlates("AAA-000", "ACA-000"), 52000)
	assert.Equal(t, countRangeBetweenLicencePlates("AAA-000", "AZA-001"), 650001)
	assert.Equal(t, countRangeBetweenLicencePlates("AAA-000", "AZA-999"), 650999)
	assert.Equal(t, countRangeBetweenLicencePlates("AAA-100", "BAA-000"), 675900)
	assert.Equal(t, countRangeBetweenLicencePlates("AAA-000", "ZAA-000"), 16900000)
	assert.Equal(t, countRangeBetweenLicencePlates("AAA-000", "ZAA-500"), 16900500)
	assert.Equal(t, countRangeBetweenLicencePlates("AAA-000", "ZZZ-999"), 17575999)
}
