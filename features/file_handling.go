package asciiart

import (
	"fmt"
	"os"
)

// SaveFile saves the given string data to a file with the specified file name.
// It returns an error if there is any issue encountered during file creation or writing.
func SaveFile(fileName string, str string) error {
	// Attempt to create the file with the provided file name
	file, err := os.Create(fileName)
	if err != nil {
		// Return an error if file creation fails
		return fmt.Errorf("failed to create file '%s': %v", fileName, err)
	}
	defer file.Close() // Close the file at the end of the function, regardless of success or failure

	// Convert the string data to bytes
	data := []byte(str)

	// Write the data to the file
	_, err = file.Write(data)
	if err != nil {
		// Return an error if writing to the file fails
		return fmt.Errorf("failed to write data to file '%s': %v", fileName, err)
	}

	// Return nil if the file writing is successful
	return nil
}
