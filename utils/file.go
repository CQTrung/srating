package utils

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)

func SanitizeFilename(fileName string) string {
	// Remove any leading or trailing spaces from the file name
	fileName = strings.TrimSpace(fileName)
	fileName = strings.ReplaceAll(fileName, "  ", " ")
	if fileName == "" {
		fileName = "default_filename"
	}

	allowedRunes := func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || r == '.' || r == '-' {
			return r
		}
		if unicode.IsSpace(r) {
			return '_'
		}
		return -1
	}

	sanitized := strings.Map(allowedRunes, fileName)

	unixTime := time.Now().Unix()
	sanitizedFileName := fmt.Sprintf("%d-%s", unixTime, sanitized)

	// Limit the filename length to 100 characters
	const maxFileNameLength = 100
	if len(sanitizedFileName) > maxFileNameLength {
		sanitizedFileName = sanitizedFileName[:maxFileNameLength]
	}

	return sanitizedFileName
}
