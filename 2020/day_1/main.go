package main

import (
	"fmt"
	"github.com/mitesh4000/AOC/utils"
	"strconv"
	"strings"
)

func main() {

	fileOutput, err := utils.ReadFile("ip_day1.txt")

	if err != nil {
		fmt.Println(err)
	}

	arrayOfStrings := strings.Split(fileOutput, "\n")
	arrayOfNumbers := make([]int, len(arrayOfStrings))

	for _, currentString := range arrayOfStrings {
		if currentString == "" {
			continue
		}
		convertedNumber, err := strconv.Atoi(currentString)
		if err != nil {
			fmt.Printf("Erro converting to number :%v", err)
		}
		fmt.Println("converted", convertedNumber)
		arrayOfNumbers = append(arrayOfNumbers, convertedNumber)
	}
	var result int
loopOtherNum:
	for _, currentNumber := range arrayOfNumbers {
		otherNumber := 2020 - currentNumber

		for _, currentOtherNumber := range arrayOfNumbers {
			if otherNumber == currentOtherNumber {
				result = otherNumber * currentNumber
				break loopOtherNum
			}
		}
	}

	fmt.Println("result: ", result)

