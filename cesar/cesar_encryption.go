package main

import (
	"fmt"
	"strings"
)

var alphabetMap map[int]string

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// EncryptUsingCesarMethod allows encrypt a word moving n times the alphabet
func EncryptUsingCesarMethod(alphabetMap map[int]string, word string, numberOfMovements int) (encryptedString string) {

	alphabetMapEncrypted := createEncryptedAlphabet(alphabetMap, numberOfMovements)
	wordSplitted := strings.Split(word, "")

	wordMap := make(map[int]string)
	for wordIndex := 0; wordIndex < len(wordSplitted); wordIndex++ {
		wordMap[wordIndex] = wordSplitted[wordIndex]
	}

	positions := []int{}
	for wordMapIndex := 0; wordMapIndex < len(wordMap); wordMapIndex++ {
		for alphabetMapEncryptedIndex := 0; alphabetMapEncryptedIndex < len(alphabetMapEncrypted); alphabetMapEncryptedIndex++ {
			if wordMap[wordMapIndex] == alphabetMapEncrypted[alphabetMapEncryptedIndex] {
				positions = append(positions, alphabetMapEncryptedIndex)
			}
		}
	}

	for positionsIndex := 0; positionsIndex < len(positions); positionsIndex++ {
		encryptedString += alphabetMap[positions[positionsIndex]]
	}

	return encryptedString
}

func createEncryptedAlphabet(alphabetMap map[int]string, numberOfMovements int) map[int]string {

	alphabetMapEncrypted := make(map[int]string)
	for alphabetMapIndex := 0; alphabetMapIndex < len(alphabetMap); alphabetMapIndex++ {

		newPosition := alphabetMapIndex + numberOfMovements
		if newPosition >= len(alphabetMap) {
			newPosition = newPosition - len(alphabetMap)
		} else if newPosition < 0 {
			newPosition = len(alphabetMap) + newPosition
		}

		alphabetMapEncrypted[newPosition] = alphabetMap[alphabetMapIndex]

	}
	return alphabetMapEncrypted
}

func createAlphabet(alphabetMap map[int]string) map[int]string {

	for i := 0; i < len(alphabet); i++ {
		alphabetMap[i] = string(alphabet[i])
	}
	return alphabetMap

}

// Alpha alphabet
func Alpha() string {
	return alphabet
}

func main() {
	alphabetMap = make(map[int]string)
	alphabetMap = createAlphabet(alphabetMap)
	fmt.Println("CASA, encrypt (1) ->", EncryptUsingCesarMethod(alphabetMap, "CASA", -1))
}
