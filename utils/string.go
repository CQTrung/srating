package utils

import (
	"strings"
)

func Concat(strs ...string) string { return strings.Join(strs, "") }
func ConcatWithSeparator(sep string, strs ...string) string {
	return strings.Join(strs, sep)
}

func ConcatWithSeparatorAndPrefix(prefix, sep string, strs ...string) string {
	return strings.Join(append([]string{prefix}, strs...), sep)
}

func ConcatWithSeparatorAndSuffix(suffix, sep string, strs ...string) string {
	return strings.Join(append(strs, suffix), sep)
}

func ConcatWithSeparatorAndPrefixAndSuffix(prefix, suffix, sep string, strs ...string) string {
	return strings.Join(append(append([]string{prefix}, strs...), suffix), sep)
}
