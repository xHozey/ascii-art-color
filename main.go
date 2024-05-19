package main

import (
	"fmt"
	"os"
	"strings"

	ft "asciiart/features"
)

func main() {
	defaultArgs := os.Args[1:]

	banner, colorFile, input, letters, outputFile := ft.ExtractOutputFlag(defaultArgs)
	if !ft.CheckValidInput(input) {
		fmt.Println("Invalid ascii characters")
		os.Exit(0)
	}

	data, err := os.ReadFile("banners/" + banner)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	splittedInput := strings.Split(input, "\\n")

	stringData := string(data[1:])
	if banner == "thinkertoy.txt" {
		stringData = strings.ReplaceAll(stringData, "\r", "")
	}
	content := strings.Split(stringData, "\n\n")
	characterMatrix := ft.ConvertToCharacterMatrix(content)
	result := ""

	if ft.OutputFlag && outputFile != "" {
		result := ft.DrawASCIIArt(characterMatrix, strings.Join(splittedInput, "\n"), letters, "")
		err := ft.SaveFile(outputFile, result)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
	if !ft.OutputFlag {

		for _, Input := range splittedInput {
			if Input == "" {
				result += "\n"
			} else {
				result += ft.DrawASCIIArt(characterMatrix, Input, letters, colorFile)
			}
		}
		fmt.Println(result)
	}
}
