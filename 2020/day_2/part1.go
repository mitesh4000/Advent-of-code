
package main

import (
	"github.com/mitesh4000/AOC/utils";
	"fmt"
	"strings"
	"strconv"
)

func Part1 (){
	var input,err = utils.ReadFile("ip_day2down.text")
	if(err!=nil){
		fmt.Println(err)
	}

	records := strings.Split(strings.TrimSpace(input),"\n")
		var validPasswords int = 0
	for _,record := range records {
		
	recordItems := strings.Split(record ," ")

	minAcur,err := strconv.Atoi(strings.Split(recordItems[0],"-")[0])

	if err!=nil {
		fmt.Println("String conversion",err)
	}
	
	maxAcur,err := strconv.Atoi(strings.Split(recordItems[0],"-")[1])
		fmt.Println("maxAcur ",maxAcur  )

	if err!=nil {
		fmt.Println("String conversion",err)
	}
	charactor := strings.Split(recordItems[1],"")[0]
	password := recordItems[2]

	numberOfChar := strings.Count(password,charactor)
	if numberOfChar <=maxAcur && numberOfChar >=minAcur {
		validPasswords +=1
	}
	}

	fmt.Println(validPasswords  )
}
