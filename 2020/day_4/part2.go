package main

import (
	"fmt"
	"strings"

	"github.com/mitesh4000/AOC/utils"
)

func Part2() {
	input, err := utils.ReadFile("ip_test.text")
	fmt.Println(err)
	if err != nil {
	}

	passPortList := strings.Split(input, "\n\n")

	type Passport struct {
		ecl    string
		pid    string
		eyr    string
		intbyr string
		iyr    string
		cid    string
		hgt    string
	}

	for index, passport := range passPortList {
		var passPortStruct Passport

		fmt.Println(passPortStruct, index)
	}

	fmt.Println(passPortList)
}
