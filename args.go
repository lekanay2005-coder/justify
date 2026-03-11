package main

import (
	"errors"
	"os"
	"strings"
)

// ParseArgs reads the command line arguments
// Returns: text, banner name, alignment, or error if invalid
func ParseArgs() (string, string, string, error) {

	args := os.Args // Get everything typed in terminal

	text := ""           // the string to convert
	banner := "standard" // default banner
	align := "left"      // default alignment

	// Only one argument: text
	if len(args) == 2 {
		text = args[1]
		return text, banner, align, nil
	}

	// Two arguments: text + banner
	if len(args) == 3 {
		text = args[1]
		banner = args[2]
		return text, banner, align, nil
	}

	// Three arguments: alignment + text + banner
	if len(args) == 4 {
		flag := args[1]  // alignment flag
		text = args[2]   // the string to convert
		banner = args[3] // banner name

		// Check if flag is correctly formatted
		if !strings.HasPrefix(flag, "--align=") {
			return "", "", "", errors.New("bad flag")
		}

		align = strings.TrimPrefix(flag, "--align=")

		// Only allow valid alignment types
		if align != "left" && align != "center" && align != "right" && align != "justify" {
			return "", "", "", errors.New("bad flag")
		}

		return text, banner, align, nil
	}

	// Wrong number of arguments
	return "", "", "", errors.New("bad args")
}