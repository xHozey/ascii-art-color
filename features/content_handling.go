package asciiart

import (
	"strings"
)

// Convert content array to a character matrix mapping ASCII characters to their line representations
func ConvertToCharacterMatrix(content []string) map[rune][]string {
	characterMatrix := map[rune][]string{}
	for i, val := range content {
		characterMatrix[rune(32+i)] = strings.Split(val, "\n")
	}
	return characterMatrix
}

// Check if there are any non-empty lines in the input lines array
func CheckEmptyLines(splittedInput []string) bool {
	for _, line := range splittedInput {
		if line != "" {
			return true
		}
	}
	return false
}

// Render the ASCII art based on the character matrix and the input lines
func DrawASCIIArtPrinted(characterMatrix map[rune][]string, Input string, letter string, ColorFile string) string {
	result := ""
	Color, Default := ColorSelection(ColorFile)
	checkLetter := CheckLettersToColor(Input, letter)

	if letter == "" {
		letter = Input
	}
	Run := false
	for n := 0; n < 8; n++ {
		for j := 0; j < len(Input); j++ {
			for k := 0; k < len(letter); k++ {
				if rune(letter[k]) == rune(Input[j]) {
					Run = true
				}
			}
			if Run && ColorFlag && checkLetter {
				result += Color + characterMatrix[rune(Input[j])][n] + Default
				Run = false
			} else if Run && ColorFlag {
				result += Color + characterMatrix[rune(Input[j])][n] + Default
			} else {
				result += characterMatrix[rune(Input[j])][n]
			}
		}
		if n < 7 {
			result += "\n"
		}
	}

	return result
}

func DrawASCIIArt(characterMatrix map[rune][]string, splittedInput []string, hasNonEmptyLines bool) string {
	result := ""
	for i, val := range splittedInput {
		if val == "" {
			if hasNonEmptyLines {
				result += "\n"
			} else if i != 0 && !hasNonEmptyLines {
				result += "\n"
			}
		} else if val != "" {
			for j := 0; j < 8; j++ {
				for _, k := range val {
					result += characterMatrix[k][j]
				}
				result += "\n"
			}
		}
	}
	return result
}
