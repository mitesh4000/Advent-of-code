package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mitesh4000/Advent-of-code/utils"
)

func main() {
	fileOut, err := utils.ReadFile("ip_day1_part2.txt")
	if err != nil {
		fmt.Println(err)
	}

	arrayOfStrings := strings.Split(fileOut, "\n")
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
mainLoop:
	for _, currentNumber := range arrayOfNumbers {
		if currentNumber < 1 {
			continue mainLoop
		}
		secNum := 2020 - currentNumber
		for _, currentOtherNumber := range arrayOfNumbers {
			if currentOtherNumber < secNum && currentOtherNumber > 0 {
				for _, thirdNum := range arrayOfNumbers {
					if currentNumber+currentOtherNumber+thirdNum == 2020 {
						fmt.Println("pl1", currentNumber)
						fmt.Println("pl2", currentOtherNumber)
						fmt.Println("pl3", thirdNum)
						result = currentNumber * currentOtherNumber * thirdNum
						break mainLoop
					} else {
						continue
					}
				}
			}
		}
	}

	fmt.Println("result: ", result)
}
