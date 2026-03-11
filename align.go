package main

import "strings"

// AlignASCII adjusts the ASCII art according to alignment type
func AlignASCII(ascii, align string) string {
	// Terminal width (fixed default for foolproof version)
	width := 80 // always safe, no panics

	// Split ASCII art into lines
	lines := strings.Split(ascii, "\n")

	for i, line := range lines {

		lineLen := len(line)

		switch align {
		case "left":
			// Left alignment = do nothing
		case "center":
			// Add spaces to center the line
			space := (width - lineLen) / 2
			if space > 0 {
				lines[i] = strings.Repeat(" ", space) + line
			}
		case "right":
			// Add spaces to move line to the right
			space := width - lineLen
			if space > 0 {
				lines[i] = strings.Repeat(" ", space) + line
			}
		case "justify":
			// Placeholder: left-align (full justify is optional)
			lines[i] = line
		}
	}

	// Join all lines back into one string
	return strings.Join(lines, "\n")
}
