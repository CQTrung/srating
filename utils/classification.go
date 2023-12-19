package utils

import "strings"

func ClassifyIncidentType(detail string, source []string) string {
	detail = strings.ToLower(detail)
	for _, s := range source {
		if strings.Contains(detail, s) {
			return s
		}
	}
	return ""
}
