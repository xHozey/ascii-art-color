package asciiart

import (
	"strings"
)

// ConvertToCharacterMatrix converts content into a character matrix.
// Each character's representation is split into lines.
func ConvertToCharacterMatrix(content []string) map[rune][]string {
	characterMatrix := map[rune][]string{}
	for i, val := range content {
		characterMatrix[rune(32+i)] = strings.Split(val, "\n")
	}
	return characterMatrix
}

// CheckValidInput checks if all characters in the input string are within the printable ASCII range.
func CheckValidInput(input string) bool {
	for _, char := range input {
		if int(char) < 32 || int(char) > 126 {
			return false
		}
	}
	return true
}

// CheckFlags checks for command line flags "--output" and "--color" in the arguments.
func CheckFlags(args []string) {
	for _, arg := range args {
		if strings.HasPrefix(arg, "--output=") {
			OutputFlag = true
		} else if strings.HasPrefix(arg, "--color=") {
			ColorFlag = true
		}
	}

}

// ValidateFileExtension validates if the file extension is ".txt".
func ValidateFileExtension(filename string) {
	if !strings.HasSuffix(filename, ".txt") || filename == "" {
		OutputUsage()
	}
}

// CheckEmptyLines checks if there are any non-empty lines in the input.
func CheckEmptyLines(splittedInput []string) bool {
	for _, line := range splittedInput {
		if line != "" {
			return true
		}
	}
	return false
}

// extractColorFlag extracts the color flag from the command line arguments.
func extractColorFlag(arg []string) string {
	color := strings.TrimPrefix(arg[0], "--color=")
	if len(color) == 0 {
		ColorUsage()
	}
	return color
}

// extractOutputFlag extracts the output flag from the command line arguments.
func extractOutputFlag(arg []string) string {
	output := strings.TrimPrefix(arg[0], "--output=")
	if len(output) == 0 {
		ColorUsage()
	}
	return output
}

// CheckLettersToColor checks if a given string contains a specific character.
func CheckLettersToColor(str string, char string) bool {
	if strings.Contains(str, char) && char != "" {
		return true
	}
	return false
}

// CheckStringToColor checks if a given string contains a specific word.
func CheckStringToColor(str string, char string) bool {
	splittedInput := strings.Split(str, " ")
	for _, val := range splittedInput {
		if strings.EqualFold(val, char) {
			return true
		}
	}
	return false
}
