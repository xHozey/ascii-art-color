package asciiart

// CheckBanner checks if the provided argument corresponds to a valid banner type.
// If it's not one of the supported banner types, sets the Banner flag to false.
func CheckBanner(arg string) {
	if arg != "standard" && arg != "thinkertoy" && arg != "shadow" {
		Banner = false
	} else {
		Banner = true
	}
}

// ApplyBanner applies the specified banner type by updating the BannerFile variable.
func applyBanner(arg string) {
	BannerFile = arg + ".txt"
}