package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/mitesh4000/AOC/utils"
)

func Part2() {
	input, err := utils.ReadFile("ip.text")
	if err != nil {
		fmt.Println(err)
	}
	passports := strings.Split(strings.TrimSpace(input), "\n\n")
	validPassCount := 0

	type field struct {
		byr int
		iyr int
		eyr int
		hgt string
		hcl int
		ecl int
		pid int
	}

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
		passportKeys := []string{}
		passportValues := []string{}
		for _, passportField := range passportSlice {
			passportKeys = append(passportKeys, strings.Split(passportField, ":")[0])
			passportValues = append(passportValues, strings.Split(passportField, ":")[1])
		}

		for _, requiredField := range fields {
			fieldExist := slices.Contains(passportKeys, requiredField)
			if !fieldExist {
				fmt.Println(index, requiredField)
				continue mainLoop
			}

			currentPassportValue := passportValues[slices.Index(passportKeys, requiredField)]
			if requiredField == "pid" {
				valid, err := regexp.MatchString(`^abc$`, currentPassportValue)
				if err != nil {
					fmt.Println(err)
				}
				if !valid {
					continue mainLoop
				}
			}
		}
		validPassCount += 1
		fmt.Println("-->", passportKeys)
	}

	fmt.Println("passportFields -->", validPassCount)
}
