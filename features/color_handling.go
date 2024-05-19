package asciiart

import (
	"fmt"
	"os"
	"strings"
)

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

func ColorSelection(s string) (string, string) {
	s = strings.ToLower(s)
	var selectedColor string
	Reset := "\033[0m"
	if ColorFlag {
		switch s {
		case "red":
			selectedColor = "\033[31m"
		case "green":
			selectedColor = "\033[32m"
		case "yellow":
			selectedColor = "\033[33m"
		case "blue":
			selectedColor = "\033[34m"
		case "purple":
			selectedColor = "\033[35m"
		case "cyan":
			selectedColor = "\033[36m"
		case "gray":
			selectedColor = "\033[90m"
		case "white":
			selectedColor = "\033[97m"
		default:
			fmt.Println("please input one of this colors")
			fmt.Println(Red + "red" + Reset + " " + Green + "green" + Reset + " " + Yellow + "yellow" + Reset + " " + Blue + "blue" + Reset + " " + Magenta + "purple" + Reset + " " + Cyan + "cyan" + Reset + " " + Gray + "gray" + Reset + " " + White + "white" + Reset)
			os.Exit(0)
		}
	}
	return selectedColor, Reset
}

func CheckLettersToColor(str string, char string) bool {
	runex := []rune(char)
	var TheresLetterToColor bool
	for _, val := range runex {
		TheresLetterToColor = strings.ContainsRune(str, val)
	}
	if !TheresLetterToColor && ColorFlag && char != "" {
		fmt.Println("invalid letters to be colored")
		os.Exit(0)
	}
	return TheresLetterToColor
}
