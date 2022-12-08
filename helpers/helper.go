package helpers

import (
	"fmt"
	"os"
)

func ReadFile(inputFile string) (string, error) {
	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	// convert the file binary into a string using string
	fileContent := string(file)

	return fileContent, err
}
