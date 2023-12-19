package utils

import "strings"

func IsImage(fileName string) bool {
	// Convert the file name to lowercase to make the comparison case-insensitive
	lowerFileName := strings.ToLower(fileName)

	// List of valid image extensions
	validExtensions := []string{".jpg", ".jpeg", ".png", ".gif"}

	for _, ext := range validExtensions {
		if strings.HasSuffix(lowerFileName, ext) {
			return true
		}
	}

	return false
}
