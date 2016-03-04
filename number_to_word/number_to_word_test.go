package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransformTensToPhrase(t *testing.T) {
	assert.Equal(t, transformTensToPhrase(1, 0), "diez")
	assert.Equal(t, transformTensToPhrase(1, 5), "quince")
	assert.Equal(t, transformTensToPhrase(2, 0), "veinte")
	assert.Equal(t, transformTensToPhrase(3, 0), "treinta")
	assert.Equal(t, transformTensToPhrase(2, 1), "veintiuno")
	assert.Equal(t, transformTensToPhrase(6, 3), "sesenta y tres")
}

func TestTransformHundredsToPhrase(t *testing.T) {
	assert.Equal(t, transformHundredsToPhrase(1, 0, 0), "cien")
	assert.Equal(t, transformHundredsToPhrase(1, 2, 3), "ciento veintitres")
	assert.Equal(t, transformHundredsToPhrase(4, 0, 0), "cuatrocientos")
	assert.Equal(t, transformHundredsToPhrase(3, 0, 2), "trescientos dos")
	assert.Equal(t, transformHundredsToPhrase(3, 4, 2), "trescientos cuarenta y dos")
	assert.Equal(t, transformHundredsToPhrase(7, 4, 9), "setecientos cuarenta y nueve")
}

func TestTransformNumberToPhrase(t *testing.T) {
	assert.Equal(t, transformNumberToPhrase(9), "nueve")
	assert.Equal(t, transformNumberToPhrase(30), "treinta")
	assert.Equal(t, transformNumberToPhrase(45), "cuarenta y cinco")
	assert.Equal(t, transformNumberToPhrase(140), "ciento cuarenta")
	assert.Equal(t, transformNumberToPhrase(168), "ciento sesenta y ocho")
	assert.Equal(t, transformNumberToPhrase(400), "cuatrocientos")
	assert.Equal(t, transformNumberToPhrase(245), "doscientos cuarenta y cinco")
	assert.Equal(t, transformNumberToPhrase(587), "quinientos ochenta y siete")
	assert.Equal(t, transformNumberToPhrase(753), "setecientos cincuenta y tres")
	assert.Equal(t, transformNumberToPhrase(409), "cuatrocientos nueve")

	assert.Equal(t, transformNumberToPhrase(2640), "dos mil seiscientos cuarenta")
	assert.Equal(t, transformNumberToPhrase(6300), "seis mil trescientos")
	assert.Equal(t, transformNumberToPhrase(999), "novecientos noventa y nueve")
	assert.Equal(t, transformNumberToPhrase(1542), "mil quinientos cuarenta y dos")
	assert.Equal(t, transformNumberToPhrase(5181), "cinco mil ciento ochenta y uno")
	assert.Equal(t, transformNumberToPhrase(53200), "cincuenta y tres mil doscientos")
	assert.Equal(t, transformNumberToPhrase(74601), "setenta y cuatro mil seiscientos uno")
	assert.Equal(t, transformNumberToPhrase(81500), "ochenta y un mil quinientos")
	assert.Equal(t, transformNumberToPhrase(97893), "noventa y siete mil ochocientos noventa y tres")
	assert.Equal(t, transformNumberToPhrase(524213), "quinientos veinticuatro mil doscientos trece")

	assert.Equal(t, transformNumberToPhrase(1000000), "un millon")
	assert.Equal(t, transformNumberToPhrase(1524213), "un millon quinientos veinticuatro mil doscientos trece")
	assert.Equal(t, transformNumberToPhrase(2524213), "dos millones quinientos veinticuatro mil doscientos trece")
	assert.Equal(t, transformNumberToPhrase(14000000), "catorce millones")
	assert.Equal(t, transformNumberToPhrase(23524213), "veintitres millones quinientos veinticuatro mil doscientos trece")
	assert.Equal(t, transformNumberToPhrase(91524213), "noventa y un millones quinientos veinticuatro mil doscientos trece")
	assert.Equal(t, transformNumberToPhrase(76524213), "setenta y seis millones quinientos veinticuatro mil doscientos trece")

	assert.Equal(t, transformNumberToPhrase(151372539), "ciento cincuenta y un millones trescientos setenta y dos mil quinientos treinta y nueve")
	assert.Equal(t, transformNumberToPhrase(100000000), "cien millones")
	assert.Equal(t, transformNumberToPhrase(999999999), "novecientos noventa y nueve millones novecientos noventa y nueve mil novecientos noventa y nueve")
}
