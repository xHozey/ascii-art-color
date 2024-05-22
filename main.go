package main

import (
	"fmt"
	"os"
	"strings"

	ft "asciiart/features"
)

func main() {
	// Retrieve command line arguments
	defaultArgs := os.Args[1:]

	// Extract flags and options from command line arguments
	colorFile, input, letters, outputFile := ft.ExtractFlag(defaultArgs)

	// Check if input contains only printable ASCII characters
	if !ft.CheckValidInput(input) {
		ft.InvalidInput()
	}

	// If output flag is provided, validate file extension
	if ft.OutputFlag {
		ft.ValidateFileExtension(outputFile)
	}

	// Read banner data from file
	data, err := os.ReadFile("banners/" + ft.BannerFile)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	// Split input by "\n" escape sequence
	splittedInput := strings.Split(input, "\\n")

	// Remove carriage return if the banner is "thinkertoy"
	stringData := string(data[1:])
	if ft.BannerFile == "thinkertoy.txt" {
		stringData = strings.ReplaceAll(stringData, "\r", "")
	}

	// Split banner content into individual characters
	content := strings.Split(stringData, "\n\n")
	characterMatrix := ft.ConvertToCharacterMatrix(content)

	// Check if input contains empty lines
	hasnewline := ft.CheckEmptyLines(splittedInput)
	result := ""

	// If output flag is provided and output file is specified
	if ft.OutputFlag && outputFile != "" {
		// Generate ASCII art and save to file
		result := ft.DrawASCIIArt(characterMatrix, splittedInput, hasnewline, letters, colorFile, input)
		err := ft.SaveFile(outputFile, result)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

	// If output flag is not provided
	if !ft.OutputFlag {
		for _, Input := range splittedInput {
			// Concatenate result for each input line
			if Input == "" {
				result += "\n"
			} else {
				result = ft.DrawASCIIArt(characterMatrix, splittedInput, hasnewline, letters, colorFile, input)
			}
		}
		fmt.Println(result)
	}
}
