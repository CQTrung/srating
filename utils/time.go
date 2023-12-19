package utils

import (
	"time"
)

func GetStartOfDayUnix() int64 {
	now := time.Now()
	location := time.FixedZone("GMT+7", 7*60*60) // Create a time zone representing GMT+7
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, location)
	return startOfDay.Unix()
}

func GetDayOfUnixTime(unixTimestamp int64) int64 {
	return ((unixTimestamp / 86400) + 3) % 7
}

func GetMonthOfUnixTime(unixTimestamp int64) int {
	t := time.Unix(unixTimestamp, 0)

	// Get the month as an integer using a switch statement
	var monthInt int
	switch t.Month() {
	case time.January:
		monthInt = 1
	case time.February:
		monthInt = 2
	case time.March:
		monthInt = 3
	case time.April:
		monthInt = 4
	case time.May:
		monthInt = 5
	case time.June:
		monthInt = 6
	case time.July:
		monthInt = 7
	case time.August:
		monthInt = 8
	case time.September:
		monthInt = 9
	case time.October:
		monthInt = 10
	case time.November:
		monthInt = 11
	case time.December:
		monthInt = 12
	}
	return monthInt
}

func GetYearOfUnixTime(unixTimestamp int64) int {
	t := time.Unix(unixTimestamp, 0)
	return t.Year()
}
