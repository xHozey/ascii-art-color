package asciiart

import (
	"fmt"
	"os"
)

// invalidBanner prints an error message for an invalid banner name and exits the program.
func invalidBanner() {
	fmt.Printf("Invalid banner name")
	os.Exit(0)
}

// invalidColor prints an error message for an invalid color and exits the program.
func invalidColor() {
	fmt.Println("please input one of these colors:")
	fmt.Println(
		Red + "red" + Reset + " " +
			Green + "green" + Reset + " " +
			Yellow + "yellow" + Reset + " " +
			Blue + "blue" + Reset + " " +
			Magenta + "purple" + Reset + " " +
			Cyan + "cyan" + Reset + " " +
			Gray + "gray" + Reset + " " +
			 Orange + "orange" + Reset + " " +
			White + "white" + Reset,
	)
	os.Exit(0)
}

// invalidLetter prints an error message for invalid letters to be colored and exits the program.
func invalidLetter() {
	fmt.Println("invalid letters to be colored")
	os.Exit(0)
}

// OutputUsage prints usage instructions for the output flag and exits the program.
func OutputUsage() {
	fmt.Fprintf(os.Stderr, "Output: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --output=<fileName.txt> something standard")
	os.Exit(0)
}

// ColorUsage prints usage instructions for the color flag and exits the program.
func ColorUsage() {
	fmt.Fprintf(os.Stderr, "Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> something")
	os.Exit(0)
}

// InvalidInput prints an error message for invalid ASCII characters and exits the program.
func InvalidInput() {
	fmt.Println("Invalid ASCII characters")
	os.Exit(0)
}

// invalidFlags prints an error message for using two flags simultaneously and exits the program.
func invalidFlags() {
	fmt.Printf("You can't use 2 flags at the same time!")
	os.Exit(0)
}
