package main

import (
	"github.com/mitesh4000/AOC/utils";
	"fmt"
	"strings"
)

func main (){
	var input,err = utils.ReadFile("ip_day3.text")
	if(err!=nil){
		fmt.Println(err)
	}
	var x,numberOfTreesEncounterd int = 0,0
	xCordinatesList := strings.Split(input,"\n")

	for index,currentXCoordinates := range xCordinatesList{
		if len(currentXCoordinates)<1{
		continue
		}
		
	x+=3

	currentXCoordinateSlice := strings.Split(currentXCoordinates,"")
	if currentXCoordinateSlice[x]=="#"{
	numberOfTreesEncounterd +=1
	}
	xNextLoop := x+3
	if len(currentXCoordinates) < xNextLoop+1 {
		xlen := len(xCordinatesList)
		totalXcoordinatesRepitetion :=  xNextLoop/xlen 
	repeatedXCoordinates := strings.Repeat(xCordinatesList[index+1], totalXcoordinatesRepitetion )
xCordinatesList[index+1] = repeatedXCoordinates  
	}
	}

	fmt.Println(numberOfTreesEncounterd )
}
