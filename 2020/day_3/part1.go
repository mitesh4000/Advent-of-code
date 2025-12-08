package main

import (
	"github.com/mitesh4000/AOC/utils";
	"fmt"
	"strings"
)

func Part1(){
	var input,err = utils.ReadFile("ip_day3.text")
	if(err!=nil){
		fmt.Println(err)
	}
	xCordinatesList := strings.Split(strings.TrimSpace(input),"\n")

	var x,numberOfTreesEncounterd int = 0,0

	
	for index,currentXCoordinates := range xCordinatesList{
		if(index == 0){
		fmt.Println("invalid at",index )
		continue
		}
		currentXCoordinateSlice := strings.Split(currentXCoordinates,"")
		currentCoordinateLength := len(currentXCoordinateSlice )
		if currentCoordinateLength <1{
		fmt.Println("invalid at",index )
		continue
		}
		
	x+=3
	

if(x>=len(currentXCoordinates)){
	x	=  x-currentCoordinateLength
}

	if currentXCoordinateSlice[x]=="#"{
		currentXCoordinateSlice[x]="x"
	numberOfTreesEncounterd +=1
	}

		currentXCoordinateSlice[x]="o"
	fmt.Println("VIEW ",currentXCoordinateSlice )
	}
	fmt.Println("RESULT ",numberOfTreesEncounterd )
}
