package asciiart

import (
	"fmt"
	"os"
)

func OptimizeINputBanner(args []string) (string, string) {
	var lettersTobeColored string

	var banner string
	if len(args) == 2 {
		banner = args[1]
		banner += ".txt"
	} else {
		banner = "standard.txt"
	}
	var input string
	if len(args) >= 1 {
		input = args[0]
	} else {
		input = ""
	}

	if !CheckValidInput(input) {
		fmt.Println("Error: The input contains characters without corresponding ASCII art representation!")
		os.Exit(1)
	}

	return banner, input
}
