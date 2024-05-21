package asciiart

import (
	"strings"
)

// ANSI color code constants
var (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[90m"
	White   = "\033[97m"
)

// ColorSelection selects the appropriate ANSI color code based on the input string.
// Returns the selected color and the ANSI reset code.
func ColorSelection(s string) (string, string) {
	// Convert the input string to lowercase for case-insensitive comparison
	s = strings.ToLower(s)
	var selectedColor string
	// Check if ColorFlag is enabled
	if ColorFlag {
		// Select color based on input string
		switch s {
		case "red":
			selectedColor = Red
		case "green":
			selectedColor = Green
		case "yellow":
			selectedColor = Yellow
		case "blue":
			selectedColor = Blue
		case "purple":
			selectedColor = Magenta
		case "cyan":
			selectedColor = Cyan
		case "gray":
			selectedColor = Gray
		case "white":
			selectedColor = White
		default:
			invalidColor()
		}
	}
	return selectedColor, Reset
}
