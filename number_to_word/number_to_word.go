package main

import (
	"strconv"
	"strings"
)

var numbersMap = map[int]string{
	111: "un", 1: "uno", 2: "dos", 3: "tres", 4: "cuatro", 5: "cinco", 6: "seis", 7: "siete",
	8: "ocho", 9: "nueve", 10: "diez", 11: "once", 12: "doce", 13: "trece", 14: "catorce",
	15: "quince", 16: "dieci", 20: "veinte", 21: "veinti", 30: "treinta", 40: "cuarenta",
	50: "cincuenta", 60: "sesenta", 70: "setenta", 80: "ochenta", 90: "noventa", 100: "cien",
	101: "cientos", 110: "ciento", 500: "quinientos", 1000: "mil", 1000000: "millon", 1000001: "millones",
}

func transformTensToPhrase(tens, units int) (resultingPhrase string) {
	if units == 0 {
		resultingPhrase = numbersMap[tens*10]
	} else if tens == 0 {
		resultingPhrase = numbersMap[units]
	} else if tens == 1 {
		resultingPhrase = numbersMap[tens*10+units]
	} else if tens == 2 {
		resultingPhrase = numbersMap[21] + numbersMap[units]
	} else {
		resultingPhrase = numbersMap[tens*10] + " y " + numbersMap[units]
	}
	return
}

func transformHundredsToPhrase(hundreds, tens, units int) (resultingPhrase string) {
	if hundreds == 1 {
		if tens == 0 && units == 0 {
			resultingPhrase = numbersMap[100]
		} else {
			resultingPhrase = numbersMap[110]
		}
	} else if hundreds == 0 {
		return
	} else if hundreds == 5 {
		resultingPhrase = numbersMap[500]
	} else if hundreds == 7 {
		resultingPhrase = "sete" + numbersMap[101]
	} else if hundreds == 9 {
		resultingPhrase = "nove" + numbersMap[101]
	} else {
		resultingPhrase = numbersMap[hundreds] + numbersMap[101]
	}
	tensPhrase := transformTensToPhrase(tens, units)
	if tensPhrase == "" {
		return
	}
	if resultingPhrase != "" {
		resultingPhrase = strings.Join([]string{resultingPhrase, tensPhrase}, " ")
	}

	return
}

func transformThousandsToPhrase(hundredsOfThousands, tensOfThousands, thousands, hundreds, tens, units int) (resultingPhrase string) {
	if hundredsOfThousands == 0 {
		if tensOfThousands == 0 && thousands == 1 {
			resultingPhrase = strings.Join([]string{numbersMap[1000], transformHundredsToPhrase(hundreds, tens, units)}, " ")
			return
		} else if thousands == 1 && tensOfThousands > 0 {
			resultingPhrase = transformTensToPhrase(tensOfThousands, 111)
		} else {
			resultingPhrase = transformTensToPhrase(tensOfThousands, thousands)
		}
	} else {
		if thousands == 1 {
			resultingPhrase = transformHundredsToPhrase(hundredsOfThousands, tensOfThousands, 111)
		} else {
			resultingPhrase = transformHundredsToPhrase(hundredsOfThousands, tensOfThousands, thousands)
		}
	}
	resultingPhrase = strings.Join([]string{resultingPhrase, numbersMap[1000], transformHundredsToPhrase(hundreds, tens, units)}, " ")
	return
}

func transformMillionsToPhrase(hundredsOfMillions, tensOfMillions, millions int) (resultingPhrase string) {
	if hundredsOfMillions == 0 {
		if tensOfMillions == 0 && millions == 1 {
			resultingPhrase = strings.Join([]string{numbersMap[111], numbersMap[1000000]}, " ")
			return
		} else if millions == 1 && tensOfMillions > 0 {
			resultingPhrase = transformTensToPhrase(tensOfMillions, 111)
		} else {
			resultingPhrase = transformTensToPhrase(tensOfMillions, millions)
		}
	} else {
		if millions == 1 {
			resultingPhrase = transformHundredsToPhrase(hundredsOfMillions, tensOfMillions, 111)
		} else {
			resultingPhrase = transformHundredsToPhrase(hundredsOfMillions, tensOfMillions, millions)
		}
	}
	if resultingPhrase != "" {
		resultingPhrase = strings.Join([]string{resultingPhrase, numbersMap[1000001]}, " ")
	}

	return
}

func transformNumberToPhrase(number int) (resultingPhrase string) {

	numberSplitted := strings.Split(strconv.Itoa(number), "")
	numberSize := len(numberSplitted)
	numbersSlice := make([]int, numberSize)

	for numberSplittedIndex := 0; numberSplittedIndex < len(numberSplitted); numberSplittedIndex++ {
		digit, _ := strconv.Atoi(numberSplitted[numberSplittedIndex])
		numbersSlice = append(numbersSlice, digit)
	}
	//Units
	if numberSize == 1 {
		resultingPhrase = numbersMap[numbersSlice[numberSize]]
		//Tens
	} else if numberSize == 2 {
		resultingPhrase = transformTensToPhrase(numbersSlice[numberSize], numbersSlice[numberSize+1])
		//Hundreds
	} else if numberSize == 3 {
		resultingPhrase = transformHundredsToPhrase(numbersSlice[numberSize], numbersSlice[numberSize+1], numbersSlice[numberSize+2])
		//Thounsands
	} else if numberSize == 4 {
		resultingPhrase = transformThousandsToPhrase(0, 0, numbersSlice[numberSize], numbersSlice[numberSize+1], numbersSlice[numberSize+2], numbersSlice[numberSize+3])
		//Tens of thousands
	} else if numberSize == 5 {
		resultingPhrase = transformThousandsToPhrase(0, numbersSlice[numberSize], numbersSlice[numberSize+1], numbersSlice[numberSize+2], numbersSlice[numberSize+3], numbersSlice[numberSize+4])
		//Hundreds of thousands
	} else if numberSize == 6 {
		resultingPhrase = transformThousandsToPhrase(numbersSlice[numberSize], numbersSlice[numberSize+1], numbersSlice[numberSize+2], numbersSlice[numberSize+3], numbersSlice[numberSize+4], numbersSlice[numberSize+5])
		//Millions
	} else if numberSize == 7 {
		if numbersSlice[numberSize+1] == 0 && numbersSlice[numberSize+2] == 0 && numbersSlice[numberSize+3] == 0 {
			hundreds := transformHundredsToPhrase(numbersSlice[numberSize+4], numbersSlice[numberSize+5], numbersSlice[numberSize+6])
			if hundreds != "" {
				resultingPhrase = transformMillionsToPhrase(0, 0, numbersSlice[numberSize]) + " " + hundreds
			} else {
				resultingPhrase = transformMillionsToPhrase(0, 0, numbersSlice[numberSize])
			}

		} else {
			resultingPhrase = transformMillionsToPhrase(0, 0, numbersSlice[numberSize]) + " " + transformThousandsToPhrase(numbersSlice[numberSize+1], numbersSlice[numberSize+2], numbersSlice[numberSize+3], numbersSlice[numberSize+4], numbersSlice[numberSize+5], numbersSlice[numberSize+6])
		}
		//Hundred of Millions
	} else if numberSize == 8 {
		if numbersSlice[numberSize+2] == 0 && numbersSlice[numberSize+3] == 0 && numbersSlice[numberSize+4] == 0 {
			hundreds := transformHundredsToPhrase(numbersSlice[numberSize+5], numbersSlice[numberSize+6], numbersSlice[numberSize+7])
			if hundreds != "" {
				resultingPhrase = transformMillionsToPhrase(0, numbersSlice[numberSize], numbersSlice[numberSize+1]) + " " + hundreds
			} else {
				resultingPhrase = transformMillionsToPhrase(0, numbersSlice[numberSize], numbersSlice[numberSize+1])
			}

		} else {
			resultingPhrase = transformMillionsToPhrase(0, numbersSlice[numberSize], numbersSlice[numberSize+1]) + " " + transformThousandsToPhrase(numbersSlice[numberSize+2], numbersSlice[numberSize+3], numbersSlice[numberSize+4], numbersSlice[numberSize+5], numbersSlice[numberSize+6], numbersSlice[numberSize+7])
		}
		//Thousands of Millions
	} else if numberSize == 9 {
		if numbersSlice[numberSize+3] == 0 && numbersSlice[numberSize+4] == 0 && numbersSlice[numberSize+5] == 0 {
			hundreds := transformHundredsToPhrase(numbersSlice[numberSize+6], numbersSlice[numberSize+7], numbersSlice[numberSize+8])
			if hundreds != "" {
				resultingPhrase = transformMillionsToPhrase(numbersSlice[numberSize], numbersSlice[numberSize+1], numbersSlice[numberSize+2]) + " " + hundreds
			} else {
				resultingPhrase = transformMillionsToPhrase(numbersSlice[numberSize], numbersSlice[numberSize+1], numbersSlice[numberSize+2])
			}
		} else {
			resultingPhrase = transformMillionsToPhrase(numbersSlice[numberSize], numbersSlice[numberSize+1], numbersSlice[numberSize+2]) + " " + transformThousandsToPhrase(numbersSlice[numberSize+3], numbersSlice[numberSize+4], numbersSlice[numberSize+5], numbersSlice[numberSize+6], numbersSlice[numberSize+7], numbersSlice[numberSize+8])
		}

	}

	return

}
