package main


import (
	"fmt"
	"os"
	"errors"
	"strings"
	"strconv"
)

func readFile(location string)(string,error){
	contents,err := os.ReadFile(location)
	if err != nil {
		return "",errors.New("error reading file")
	}
	return string(contents),nil
}


func main() {
   fileOutput,err := readFile("aoc_input_2020_1.txt")
	 
	 if err != nil {
	 fmt.Println(err)
	 }

	 arrayOfStrings := strings.Split(fileOutput,"\n")
	 arrayOfNumbers := make([]int,len(arrayOfStrings))


	 for _,currentString := range arrayOfStrings   {
		 if currentString == "" {
			 continue
		 }
		 convertedNumber,err := strconv.Atoi(currentString)
		 if err!=nil{
			 fmt.Printf("Erro converting to number :%v",err)
		 }
fmt.Println("converted",convertedNumber)
		 arrayOfNumbers= append(arrayOfNumbers ,convertedNumber )	 
	 }
	var result int 
	loopOtherNum : for _,currentNumber := range arrayOfNumbers {
		 otherNumber := 2020 - currentNumber

		  for _,currentOtherNumber := range arrayOfNumbers {
		 if otherNumber == currentOtherNumber {
result = otherNumber * currentNumber
break loopOtherNum
		 }
	 }
	 }
	 
	 fmt.Println("result: ",result)
}
