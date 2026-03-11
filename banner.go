package main

import (
	"errors"
	"os"
	"strings"
)

// GenerateASCII reads the banner file and builds ASCII art
func GenerateASCII(text, banner string) (string, error) {

	// Open the banner file like "standard.txt"
	data, err := os.ReadFile(banner + ".txt")
	if err != nil {
		return "", errors.New("banner file not found")
	}

	// Split the file into individual lines
	lines := strings.Split(string(data), "\n")

	result := "" // will store the final ASCII art

	// Each ASCII letter uses 8 rows
	for i := 0; i < 8; i++ {
		// Loop through every character in the text
		for _, char := range text {
			// Find the start line of this character in banner file
			index := (int(char) - 32) * 9 // 8 lines + 1 empty line per char

			// Add this row of the character to the result
			result += lines[index+i]
		}

		// Move to next line after finishing one row
		result += "\n"
	}

	return result, nil
}