package utils

import (
	"crypto/rand"
	"strconv"
	"strings"
)

func GenerateRandomString(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	chars := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charsLen := len(chars)

	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = chars[int(randomBytes[i])%charsLen]
	}

	return string(result), nil
}

func GenerateIncidentCode(code string) string {
	codeNumber := code[3:]
	codeNumberInt, err := strconv.Atoi(codeNumber)
	if err != nil {
		return ""
	}
	codeNumberInt++
	codeNumber = strconv.Itoa(codeNumberInt)
	padding := 6 - len(codeNumber)
	code = "TTD" + strings.Repeat("0", padding) + codeNumber
	return code
}
