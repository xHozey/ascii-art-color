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
	checkLetter := ft.CheckLettersToColor(input, letters)
	if letters != "" && !checkLetter {
		fmt.Println("Invalid Letters To be Color")
		os.Exit(0)
	}
	result := ""
	for _, Input := range splittedInput {
		if !ft.CheckValidInput(Input) {
			fmt.Println("gwdfhopmligukfjhfwchnb hgdc,vkhv vjhv nbyu b ;jh n yj ")
			os.Exit(0)
		} else if Input == "" {
			result += "\n"
		} else {
			result += ft.DrawASCIIArt(characterMatrix, Input, letters, colorFile)
		}
	}
	fmt.Println(result)
	if outputFile != "" {
		result := ft.DrawASCIIArt(characterMatrix, strings.Join(splittedInput, "\n"), letters, "")
		err := ft.SaveFile(outputFile, result)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
