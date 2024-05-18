package asciiart

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var OutputFlag, ColorFlag bool

func ExtractOutputFlag(args []string) (string, string, []string) {
	var outputFile string
	var colorFile string
	var filteredArgs []string

	for _, val := range args {
		if strings.HasPrefix(val, "--output=") {
			OutputFlag = true
		}
		if strings.HasPrefix(val, "--color=") {
			ColorFlag = true
		}
		if ColorFlag && OutputFlag {
			fmt.Println("you can use only one flag!")
			os.Exit(0)
		}
	}

	// check if arguments are given as one string
	if len(args) == 1 {
		tempArgs := strings.Fields(args[0])
		for _, v := range tempArgs {
			if strings.HasPrefix(v, "--") {
				args = tempArgs
			}
		}
	}

	for i, arg := range args {
		if strings.HasPrefix(arg, "--output=") && i == 0 {
			outputFile = strings.TrimPrefix(arg, "--output=")
			if len(outputFile) == 0 {
				OutputUsage()
				os.Exit(0)
			}
		} else if strings.HasPrefix(arg, "--color=") && i == 0 {
			colorFile = strings.TrimPrefix(arg, "--color=")
			outputFile = ""

			if len(colorFile) == 0 {
				OutputUsage()
				os.Exit(0)
			}
		} else if strings.HasPrefix(arg, "--") {
			OutputUsage()
			os.Exit(0)
		} else {
			filteredArgs = append(filteredArgs, arg)
		}
	}

	if err := CheckArguments(filteredArgs); err {
		OutputUsage()
		os.Exit(0)
	}
	if outputFile != "" {
		if err := ValidateFileExtension(outputFile); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(0)
		}
	}

	return outputFile, colorFile, filteredArgs
}

func CheckArguments(args []string) bool {
	argCount := len(args)
	if !ColorFlag {
		if argCount == 2 {
			if args[1] != "shadow" && args[1] != "standard" && args[1] != "thinkertoy" {
				fmt.Printf("invalid banner type '%s'\nAvailable banner types are: 'standard' (default), 'shadow', and 'thinkertoy'\n", args[1])
				os.Exit(0)
			}
		}
		if argCount > 2 {
			return true
		}
	} else {
		if argCount == 3 {
			if args[2] != "shadow" && args[2] != "standard" && args[2] != "thinkertoy" {
				fmt.Printf("invalid banner type '%s'\nAvailable banner types are: 'standard' (default), 'shadow', and 'thinkertoy'\n", args[1])
				os.Exit(0)
			}
		}
	}
	return false
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
}

func ColorUsage() {
	fmt.Fprintf(os.Stderr, "\n Usage: go run . [OPTION] [STRING]\n\n EX: go run . --color=<color> <letters to be colored> something\n\n")
}
