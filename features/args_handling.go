package asciiart

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var OutputFlag, ColorFlag bool

func ExtractOutputFlag(args []string) (string, string, string, string, string) {
	var (
		outputFile     string
		colorFile      string
		banner         = "standard.txt"
		input, letters string
	)
	if len(args) > 4 {
		fmt.Println("Too many arguments")
		os.Exit(0)
	}

	switch len(args) {
	case 4:
		if strings.HasPrefix(args[0], "--color=") {
			colorFile = strings.TrimPrefix(args[0], "--color=")
			if len(colorFile) == 0 {
				ColorUsage()
			}
			letters = args[1]
			input = args[2]
			banner = args[3] + ".txt"
			ColorFlag = true
		}
		if strings.HasPrefix(args[0], "--output=") {
			OutputUsage()
		}
	case 3:
		if strings.HasPrefix(args[0], "--color=") {
			colorFile = strings.TrimPrefix(args[0], "--color=")
			if len(colorFile) == 0 {
				ColorUsage()
			}
			if args[2] == "standard" || args[2] == "thinkertoy" || args[2] == "shadow" {
				input = args[1]
				banner = args[2] + ".txt"
			} else {
				letters = args[1]
				input = args[2]
			}
			ColorFlag = true
		} else if strings.HasPrefix(args[0], "--output=") {

			outputFile = strings.TrimPrefix(args[0], "--output=")
			if len(outputFile) == 0 {
				OutputUsage()
			}
			input = args[1]
			if args[2] != "standard" && args[2] != "thinkertoy" && args[2] != "shadow" {
				fmt.Println("invalid Banner Name")
				os.Exit(0)
			} else {
				banner = args[2] + ".txt"
			}
			OutputFlag = true
		} else {
			ColorUsage()
		}
	case 2:
		if strings.HasPrefix(args[0], "--color=") {

			colorFile = strings.TrimPrefix(args[0], "--color=")
			if len(colorFile) == 0 {
				ColorUsage()
			}
			input = args[1]
			ColorFlag = true
		} else if strings.HasPrefix(args[0], "--output=") {

			outputFile = strings.TrimPrefix(args[0], "--output=")
			if len(outputFile) == 0 {
				OutputUsage()
			}
			input = args[1]
			OutputFlag = true
		} else if args[1] == "standard" || args[1] == "thinkertoy" || args[1] == "shadow" {
			input = args[0]
			banner = args[1] + ".txt"
		} else {
			fmt.Println("Invalid bannerfile")
			os.Exit(0)
		}
	case 1:
		input = args[0]
	case 0:
		ColorUsage()
	}
	// check if arguments are given as one string

	return banner, colorFile, input, letters, outputFile
}

// validates if the input contains only printable ASCII characters
func CheckValidInput(input string) bool {
	for _, char := range input {
		if int(char) < 32 || int(char) > 126 {
			return false
		}
	}
	return true
}

func ValidateFileExtension(filename string) error {
	acceptableExtensions := []string{".txt", ".json"}
	extension := strings.ToLower(filepath.Ext(filename))
	if extension == "" {
		return fmt.Errorf("please use one of the following extensions for the output file: .txt")
	}
	for _, ext := range acceptableExtensions {
		if extension == ext {
			return nil
		}
	}
	return fmt.Errorf("invalid file extension '%s' for --output option. Please use one of the following extensions: .txt", extension)
}

func OutputUsage() {
	fmt.Fprintf(os.Stderr, "\n   Output: go run . [OPTION] [STRING] [BANNER]\n\n   Example: go run . --output=<fileName.txt> something standard\n\n")
	os.Exit(0)
}

func ColorUsage() {
	fmt.Fprintf(os.Stderr, "\n Usage: go run . [OPTION] [STRING]\n\n EX: go run . --color=<color> <letters to be colored> something\n\n")
	os.Exit(0)
}
