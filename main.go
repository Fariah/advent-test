package main

import (
	"advent-test/tasks"
	"fmt"
	"os"
)

func main() {
	taskNumber := os.Args[1:]

	switch taskNumber[0] {
	case "01":
		tasks.Execute01()
	case "02":
		tasks.Execute02()
	default:
		fmt.Printf("please, input number beetwen 01 and 02")
	}
}
