package main

import (
	"github.com/mitesh4000/AOC/utils";
	"fmt"
	"strings"
)



func Part2(){
	var input,err = utils.ReadFile("ip_day3.text")
	if(err!=nil){
		fmt.Println(err)
	}
	xCordinatesList := strings.Split(strings.TrimSpace(input),"\n")

	
	type slop struct {
		Right int
		down int
	} 
	var slops = [5]slop{
		{Right: 1, down: 1},
		{Right: 3, down: 1}, 
		{Right: 5, down: 1},
		{Right: 7, down: 1},
		{Right: 1, down: 2},
	}

	var slopResults []int
	for _,currentSlop := range slops{

	var y,x,numberOfTreesEncounterd int =0, 0,0
	yCordinatesLength := len(xCordinatesList) 

for {
	y+=currentSlop.down
	x+=currentSlop.Right	


	if y>= yCordinatesLength {
		break
	}
	currentXCoordinates := xCordinatesList[y] 
		currentXCoordinateSlice := strings.Split(currentXCoordinates,"")
		currentCoordinateLength := len(currentXCoordinateSlice )
		if currentCoordinateLength <1{
		fmt.Println("invalid at",y )
		continue
		}
		
	

if(x>=len(currentXCoordinates)){
	x	=  x-currentCoordinateLength
}

	if currentXCoordinateSlice[x]=="#"{
		currentXCoordinateSlice[x]="x"
	numberOfTreesEncounterd +=1
	}
		currentXCoordinateSlice[x]="o"
	}


	fmt.Println("RESULT ",numberOfTreesEncounterd )
	slopResults=append(slopResults,numberOfTreesEncounterd )
	numberOfTreesEncounterd =0

	}

	fmt.Println("SLOP RESULT ",slopResults )
	answer := 1
	for i:=0;i<len(slopResults);i++{
		answer = answer*slopResults[i]
	}

	fmt.Println("ANS",answer )

	}


