package asciiart

import (
	"fmt"
	"os"
	"strings"
)

func CheckLettersToColor(str string, char string) bool {
	TheresLetterToColor := strings.Contains(str, char)
	if !TheresLetterToColor {
		fmt.Println("invalid letters to be colored")
		os.Exit(0)
	}
	return TheresLetterToColor
}
