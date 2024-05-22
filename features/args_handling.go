package asciiart

var OutputFlag, ColorFlag, Banner bool
var BannerFile string = "standard.txt"

// ExtractFlag extracts flags and arguments from the command line.
func ExtractFlag(args []string) (string, string, string, string) {
	var (
		outputFile string
		colorFile  string
		input      string
		letters    string
	)

	// Check for specified flags and handle accordingly
	CheckFlags(args)

	// Handling cases when both color and output flags are specified
	if ColorFlag && OutputFlag {
		invalidFlags()
	}

	// Handling cases when there are more than 4 arguments provided
	if len(args) > 4 {
		ColorUsage()
	}

	// Switching based on the number of arguments provided
	switch len(args) {
	case 4:
		// Check if banner flag is specified
		CheckBanner(args[3])
		if ColorFlag {
			colorFile = extractColorFlag(args)
			letters = args[1]
			input = args[2]
			BannerFile = args[3] + ".txt"
		} else {
			ColorUsage()
		}
		if OutputFlag {
			OutputUsage()
		}
		if Banner {
			applyBanner(args[3])
		} else {
			invalidBanner()
		}
	case 3:
		// Check if banner flag is specified
		CheckBanner(args[2])
		if ColorFlag {
			colorFile = extractColorFlag(args)
			letters = args[1]
			input = args[2]
		} else if OutputFlag {
			outputFile = extractOutputFlag(args)
			input = args[1]
			if Banner {
				applyBanner(args[2])
			} else {
				invalidBanner()
			}
		} else {
			ColorUsage()
		}
	case 2:
		// Check if banner flag is specified
		CheckBanner(args[1])
		if ColorFlag {
			colorFile = extractColorFlag(args)
			input = args[1]
		} 
		if OutputFlag {
			outputFile = extractOutputFlag(args)
			input = args[1]
		}
		if Banner {
			applyBanner(args[1])
			input = args[0]
		} else if !Banner && !OutputFlag && !ColorFlag {
			invalidBanner()
		}
	case 1:
		// Handling case when only one argument is provided
		if ColorFlag {
			ColorUsage()
		} else if OutputFlag {
			OutputUsage()
		}
		input = args[0]
	case 0:
		// Handling case when no arguments are provided
		ColorUsage()
	}

	return colorFile, input, letters, outputFile
}
