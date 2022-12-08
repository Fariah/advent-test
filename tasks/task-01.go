package tasks

import (
	"advent-test/helpers"
	"fmt"
	"strconv"
	"strings"
)

func Execute01() {
	fmt.Printf("Start task 01\n")
	content, _ := helpers.ReadFile("inputs/input-01.txt")
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
