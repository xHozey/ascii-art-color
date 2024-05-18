package main

import (
	"fmt"
	"os"
	"strings"

	ft "asciiart/features"
)

func main() {
	defaultArgs := os.Args[1:]
	if len(defaultArgs) == 0 || len(defaultArgs[0]) == 0 {
		return
	}

	outputFile, colorFile, args := ft.ExtractOutputFlag(defaultArgs)

	banner, input, letters := ft.OptimizeInputBanner(args)

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

	result := ft.DrawASCIIArt(characterMatrix, splittedInput)
	checkLetter := ft.CheckLettersToColor(input, letters)

	if outputFile != "" {
		err := ft.SaveFile(outputFile, result)
		if err != nil {
			fmt.Println("Error:", err)
		}
	} else if ft.ColorFlag && letters == "" {
		color, Reset := ft.ColorSelection(colorFile)
		fmt.Printf(color + result + Reset)
	} else if ft.ColorFlag && letters != "" {
		if checkLetter {
			color, Reset := ft.ColorSelection(colorFile)
			fmt.Printf(color + result + Reset)
		} else {
			fmt.Println("invalid letters to be colored")
		}
	} else {
		fmt.Printf(result)
	}
}
