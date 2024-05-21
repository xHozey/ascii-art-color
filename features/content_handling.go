package asciiart

import "strings"

// DrawASCIIArt generates ASCII art from a character matrix and input text.
// It applies colors if ColorFlag is enabled and applies letter-specific coloring if specified.
func DrawASCIIArt(characterMatrix map[rune][]string, splittedInput []string, hasNonEmptyLines bool, letter string, ColorFile string, input string) string {
	result := ""
	Run := false
	Color, Default := ColorSelection(ColorFile)

	// Check if specific letters are designated for coloring
	letterCheck := false
	stringCheck := false
	if len(letter) == 1 {
		letterCheck = CheckLettersToColor(input, letter)
	} else {
		stringCheck = CheckStringToColor(input, letter)
	}

	if letter == "" {
		letter = input
	}

	for i, line := range splittedInput {
		if line == "" {
			// If the line is empty, add a newline character to the result
			if hasNonEmptyLines {
				result += "\n"
			} else if i != 0 && !hasNonEmptyLines {
				result += "\n"
			}
		} else {
			if !ColorFlag && !letterCheck && !stringCheck {
				// If ColorFlag is not enabled, draw ASCII art without color
				for j := 0; j < 8; j++ {
					for _, k := range line {
						result += characterMatrix[k][j]
					}
					result += "\n"
				}
			} else {
				for j := 0; j < 8; j++ {
					splitedLine := strings.Split(line, " ")
					for _, word := range splitedLine {
						// If ColorFlag is enabled, apply coloring
						for _, k := range word {
							// Check if the current character should be colored
							if letterCheck {
								for l := 0; l < len(letter); l++ {
									if rune(letter[l]) == rune(k) {
										Run = true
									}
								}
							}
							if stringCheck && word == letter {
								Run = true
							}
							if Run && ColorFlag && letterCheck && !stringCheck {
								// Apply color to the character
								result += Color + characterMatrix[k][j] + Default
								Run = false
							} else if Run && ColorFlag && !letterCheck {
								// Apply color to the character
								result += Color + characterMatrix[k][j] + Default
								Run = false
							} else {
								result += characterMatrix[k][j]
							}
						}
					}
					result += "\n"
				}
			}
		}
	}
	return result
}
