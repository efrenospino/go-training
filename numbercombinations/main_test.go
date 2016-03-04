package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeCombinationsMap(t *testing.T) {
	_, numberOfCombinations := makeCombinationsMap(2, 5)
	fmt.Println("Number Of Combinations:", numberOfCombinations)
	assert.Equal(t, numberOfCombinations, 15)
	_, numberOfCombinations = makeCombinationsMap(2, 6)
	fmt.Println("Number Of Combinations:", numberOfCombinations)
	assert.Equal(t, numberOfCombinations, 23)
	_, numberOfCombinations = makeCombinationsMap(2, 10)
	fmt.Println("Number Of Combinations:", numberOfCombinations)
	assert.Equal(t, numberOfCombinations, 69)
	_, numberOfCombinations = makeCombinationsMap(2, 100)
	fmt.Println("Number Of Combinations:", numberOfCombinations)
	//assert.Equal(t, numberOfCombinations, 69)
}

func TestOrderingCombinationsMap(t *testing.T) {
	combinations, _ := makeCombinationsMap(2, 5)
	combinationsOrdered, indexes := orderingCombinationsMap(combinations)
	assert.Equal(t, combinationsOrdered[indexes[0]], 22)
}

/*
 22 23 &24 25 &26 27 28 29 210
 32 33 34 35 36 37 38 39 310
 &42 &43 44 45 46 47 48 49 410
 52 53 54 55 56 57 58 59 510
 62 63 64 65 66 67 68 69 610
 72 73 74 75 76 77 78 79 710
 82 83 84 85 86 87 88 89 810
 92 93 94 95 96 97 98 99 910
 102 103 104 105 106 107 108 109 110
*/
