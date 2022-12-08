package main

import (
	"advent-test/helpers"
	"fmt"
	"strconv"
	"strings"
)

const INPUT = "inputs/input-01.txt"

func main() {
	fmt.Printf("Start task 01\n")
	content, _ := helpers.ReadFile(INPUT)
	blocks := strings.Split(content, "\n")
	calculation := 0
	biggestCalories := 0

	for _, element := range blocks {
		call, _ := strconv.Atoi(element)

		if element == "" {
			if calculation >= biggestCalories {
				biggestCalories = calculation
			}
			calculation = 0
		} else {
			calculation = calculation + call
		}
	}

	fmt.Printf("Biggest: %v\n", biggestCalories)
}
