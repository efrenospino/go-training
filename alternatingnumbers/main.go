package main

import "fmt"

func deleteLeftToRigth(numbers []int) []int {
	for index := 0; index < len(numbers); index++ {
		numbers[index] = 0
		index++
	}
	numbersResult := []int{}
	for index := 0; index < len(numbers); index++ {
		if numbers[index] > 0 {
			numbersResult = append(numbersResult, numbers[index])
		}
	}
	return numbersResult
}

func deleteRigthToLeft(numbers []int) []int {
	for index := len(numbers) - 1; index >= 0; index-- {
		numbers[index] = 0
		index--
	}
	numbersResult := []int{}
	for index := 0; index < len(numbers); index++ {
		if numbers[index] > 0 {
			numbersResult = append(numbersResult, numbers[index])
		}
	}
	return numbersResult
}

func alternateNumbers(n int) {
	var numbers []int
	for index := 1; index <= n; index++ {
		numbers = append(numbers, index)
	}
	fmt.Println(numbers)
	for len(numbers) > 0 {
		numbers = deleteLeftToRigth(numbers)
		fmt.Println(numbers)
		numbers = deleteRigthToLeft(numbers)
		fmt.Println(numbers)
	}
}

func main() {
	alternateNumbers(15)
}
