package main

import "fmt"

// main is the first function that runs when you start the program
func main() {

	// Read arguments from terminal
	// Returns: text, banner, alignment type, or error if wrong input
	text, banner, align, err := ParseArgs()

	// If arguments are wrong, show usage and stop
	if err != nil {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		return
	}

	// Generate ASCII art from the text using the banner file
	ascii, err := GenerateASCII(text, banner)

	// If the banner file cannot be read, print error and stop
	if err != nil {
		fmt.Println(err)
		return
	}

	// Align the ASCII art according to the alignment flag
	final := AlignASCII(ascii, align)

	// Print the final ASCII art to the terminal
	fmt.Print(final)
}