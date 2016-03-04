package main

import (
	"math"
	"strconv"
	"strings"
)

var alphabetMap map[int]string

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func createAlphabet(alphabetMap map[int]string) map[int]string {
	alphabetMap = make(map[int]string)
	for i := 0; i < len(alphabet); i++ {
		alphabetMap[i] = string(alphabet[i])
	}
	return alphabetMap
}

func getIndexOfLetter(letter string, alphabetMap map[int]string) (index int) {
	index = -1
	for alphabetIndex := 0; alphabetIndex < len(alphabetMap); alphabetIndex++ {
		if letter == alphabetMap[alphabetIndex] {
			index = alphabetIndex
		}
	}
	return
}

func countRangeBetweenNumbers(numbersInLicensePlate1, numbersInLicensePlate2 int) int {
	return int(math.Abs(float64(numbersInLicensePlate1) - float64(numbersInLicensePlate2)))
}

func countRangeBetweenLetters(lettersInLicensePlate1, lettersInLicensePlate2 string) int {

	alphabetMap := createAlphabet(alphabetMap)

	lettersArray1 := strings.Split(lettersInLicensePlate1, "")
	lettersArray2 := strings.Split(lettersInLicensePlate2, "")

	letters1Indexes := []int{}
	letters2Indexes := []int{}

	for letterArrayIndex := 0; letterArrayIndex < len(lettersArray1); letterArrayIndex++ {
		letters1Indexes = append(letters1Indexes, getIndexOfLetter(lettersArray1[letterArrayIndex], alphabetMap))
		letters2Indexes = append(letters2Indexes, getIndexOfLetter(lettersArray2[letterArrayIndex], alphabetMap))
	}

	counterFirstLetter := math.Abs(float64(letters1Indexes[0]) - float64(letters2Indexes[0]))
	counterSecondLetter := math.Abs(float64(letters1Indexes[1]) - float64(letters2Indexes[1]))
	counterThirdLetter := math.Abs(float64(letters1Indexes[2]) - float64(letters2Indexes[2]))

	counter := 0

	counter += int(counterFirstLetter) * (len(alphabet) * len(alphabet))
	counter += int(counterSecondLetter) * len(alphabet)
	counter += int(counterThirdLetter)

	return counter
}

func countRangeBetweenLicencePlates(licensePlate1, licensePlate2 string) int {

	licensePlate1Splitted := strings.Split(licensePlate1, "-")
	licensePlate2Splitted := strings.Split(licensePlate2, "-")

	rangeBetweenLetters := countRangeBetweenLetters(licensePlate1Splitted[0], licensePlate2Splitted[0])
	numbers1, _ := strconv.Atoi(licensePlate1Splitted[1])
	numbers2, _ := strconv.Atoi(licensePlate2Splitted[1])

	rangeBetweenNumbers := countRangeBetweenNumbers(numbers1, numbers2)
	rangeBetweenLicensePlates := 0

	if countRangeBetweenLetters("AAA", licensePlate1Splitted[0]) > countRangeBetweenLetters("AAA", licensePlate2Splitted[0]) {
		if numbers1 > numbers2 {
			rangeBetweenLicensePlates = rangeBetweenLetters*1000 + rangeBetweenNumbers
		} else {
			rangeBetweenLicensePlates = rangeBetweenLetters*1000 - rangeBetweenNumbers
		}
	} else {
		if numbers1 > numbers2 {
			rangeBetweenLicensePlates = rangeBetweenLetters*1000 - rangeBetweenNumbers
		} else {
			rangeBetweenLicensePlates = rangeBetweenLetters*1000 + rangeBetweenNumbers
		}
	}

	return rangeBetweenLicensePlates
}
