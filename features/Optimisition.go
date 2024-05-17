package asciiart

import (
	"fmt"
	"os"
)

func OptimizeINputBanner(args []string) (string, string, string) {
	var banner string
	var letters string
	if !ColorFlag {
		if len(args) == 2 {
			banner = args[1]
			banner += ".txt"
		} else {
			banner = "standard.txt"
		}
	} else {
		if len(args) == 3 {
			banner = args[2]
			banner += ".txt"
		} else {
			banner = "standard.txt"
		}
	}
	var input string
	if !ColorFlag {
		if len(args) >= 1 {
			input = args[0]
		} else {
			input = ""
		}
	} else {
		letters = args[0]
		if len(args) >= 1 {
			input = args[1]
		} else {
			input = ""
		}
	}

	if !CheckValidInput(input) {
		fmt.Println("Error: The input contains characters without corresponding ASCII art representation!")
		os.Exit(1)
	}

	return banner, input, letters
}
