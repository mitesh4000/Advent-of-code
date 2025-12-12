package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mitesh4000/AOC/utils"
)

func Part1() {
	input, err := utils.ReadFile("ip.text")
	if err != nil {
		fmt.Println(err)
	}
	passports := strings.Split(strings.TrimSpace(input), "\n\n")
	validPassCount := 0

	fields := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}
	fmt.Println(fields)

mainLoop:
	for index, passport := range passports {
		passport = strings.ReplaceAll(passport, "\n", " ")
		passportSlice := strings.Split(passport, " ")

		if len(passportSlice) < len(fields) {
			continue mainLoop
		}
		passportFields := []string{}
		for _, passportField := range passportSlice {
			field := strings.Split(passportField, ":")[0]
			passportFields = append(passportFields, field)
		}

		for _, requiredField := range fields {
			fieldExist := slices.Contains(passportFields, requiredField)
			if !fieldExist {
				fmt.Println(index, requiredField)
				continue mainLoop
			}
		}
		validPassCount += 1
		fmt.Println("-->", passportFields)
	}

	fmt.Println("passportFields -->", validPassCount)
}
