package generatecolor

import (
	"strconv"
	"strings"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Person is a type person
type Person struct {
	Name  string
	Color string
}

func GenerateLetterValues() map[string]int {
	colorsMap := make(map[string]int)
	for alphabetIndex := 0; alphabetIndex < len(alphabet); alphabetIndex++ {
		byteColor := []byte(string(alphabet[alphabetIndex]))
		intColor := (26 * int(byteColor[0])) % 255
		colorsMap[string(alphabet[alphabetIndex])] = intColor
	}
	return colorsMap
}

func GenerateValuesByLettersInName(name string) string {
	nameSplitted := strings.Split(strings.ToUpper(name), "")
	colorsMap := GenerateLetterValues()
	valuesToIn := []string{}
	for nameSplittedIndex := 0; nameSplittedIndex < len(nameSplitted); nameSplittedIndex++ {
		in := colorsMap[nameSplitted[nameSplittedIndex]]
		valuesToIn = append(valuesToIn, strconv.Itoa(in))
	}
	intColorValue, _ := strconv.Atoi(strings.Join(valuesToIn, ""))
	hexColorValue := "#" + strconv.FormatInt(int64(intColorValue%255), 16) + strconv.FormatInt(int64(3*intColorValue%255), 16) + strconv.FormatInt(int64(5*intColorValue%255), 16)
	return hexColorValue
}
