package main

import (
	"fmt"
	"math"
	"sort"
)

func valueExists(combinationsMap map[int]int, value int) (ans bool) {
	for _, valueInMap := range combinationsMap {
		if value == valueInMap {
			ans = true
		}
	}
	return
}

func makeCombinationsMap(lowerLimit, upperLimit int) (combinationsMap map[int]int, numberOfCombinations int) {

	combinationsMap = make(map[int]int)

	for a := lowerLimit; a <= upperLimit; a++ {
		for b := lowerLimit; b <= upperLimit; b++ {
			powValue := int(math.Pow(float64(a), float64(b)))
			if !valueExists(combinationsMap, powValue) {
				combinationsMap[a*10+b] = powValue
				numberOfCombinations++
			}
		}
	}
	return
}

func orderingCombinationsMap(combinationsMap map[int]int) (combinationsMapOrdered map[int]int, indexes []int) {

	combinationsMapOrdered = make(map[int]int)
	for k, v := range combinationsMap {
		combinationsMapOrdered[v] = k
		indexes = append(indexes, v)
	}
	sort.Ints(indexes)
	for _, v := range indexes {
		fmt.Printf("%d = %d\n", combinationsMapOrdered[v], v)
	}

	//fmt.Println(indexes)

	return
}
