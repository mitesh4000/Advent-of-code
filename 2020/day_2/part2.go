
package main

import (
	"github.com/mitesh4000/AOC/utils";
	"fmt"
	"strings"
	"strconv"
)

func Part2(){
	var input,err = utils.ReadFile("ip_day2down.text")
	if(err!=nil){
		fmt.Println(err)
	}

	records := strings.Split(strings.TrimSpace(input),"\n")
		var validPasswords int = 0
	for _,record := range records {
		
	recordItems := strings.Split(record ," ")

	minAcur,err := strconv.Atoi(strings.Split(recordItems[0],"-")[0])
		fmt.Println("minAcur ",minAcur )

	if err!=nil {
		fmt.Println("String conversion",err)
	}
	
	maxAcur,err := strconv.Atoi(strings.Split(recordItems[0],"-")[1])

	if err!=nil {

		fmt.Println("String conversion",err)
	}
	charactor := recordItems[1][0]
	password := recordItems[2]

	
	if password[minAcur-1] == charactor {

		validPasswords +=1
	if password[maxAcur-1] == charactor {
		
		validPasswords -=1
	}

	} else if password[maxAcur-1] == charactor {

		validPasswords +=1
	}

	}

	fmt.Println(validPasswords  )
}
