package timer

import (
	"time"
)

type Now struct {
	time.Time
	TimeFormats []string
}

func (now *Now) IsMorning() bool {
	return now.Unix() < now.BeginningOfMidDay().Unix()
}

func (now *Now) BeginningOfMinute() time.Time {
	return now.Truncate(time.Minute)
}

func (now *Now) BeginningOfHour() time.Time {
	return now.Truncate(time.Hour)
}

func (now *Now) BeginningOfDay() time.Time {
	d := time.Duration(-now.Hour()) * time.Hour
	return now.BeginningOfHour().Add(d)
}

func (now *Now) BeginningOfMidDay() time.Time {
	d := time.Duration(12) * time.Hour
	return now.BeginningOfDay().Add(d)
}

func (now *Now) BeginningOfMonth() time.Time {
	t := now.BeginningOfDay()
	d := time.Duration(-(t.Day())+1) * 24 * time.Hour
	return t.Add(d)
}

func (now *Now) BeginningOfQuarter() time.Time {
	month := now.BeginningOfMonth()
	offset := (int(month.Month()) - 1) % 3
	return month.AddDate(0, -offset, 0)
}

func (now *Now) BeginningOfYear() time.Time {
	t := now.BeginningOfDay()
	d := time.Duration(-(t.YearDay())+1) * 24 * time.Hour
	return t.Truncate(time.Hour).Add(d)
}

func (now *Now) EndOfMinute() time.Time {
	return now.BeginningOfMinute().Add(time.Minute - time.Nanosecond)
}

func (now *Now) EndOfHour() time.Time {
	return now.BeginningOfHour().Add(time.Hour - time.Nanosecond)
}

func (now *Now) EndOfDay() time.Time {
	return now.BeginningOfDay().Add(24*time.Hour - time.Nanosecond)
}

func (now *Now) EndOfMonth() time.Time {
	return now.BeginningOfMonth().AddDate(0, 1, 0).Add(-time.Nanosecond)
}

func (now *Now) EndOfQuarter() time.Time {
	return now.BeginningOfQuarter().AddDate(0, 3, 0).Add(-time.Nanosecond)
}

func (now *Now) EndOfYear() time.Time {
	return now.BeginningOfYear().AddDate(1, 0, 0).Add(-time.Nanosecond)
}

func (now *Now) Monday() time.Time {
	t := now.BeginningOfDay()
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	d := time.Duration(-weekday+1) * 24 * time.Hour
	return t.Truncate(time.Hour).Add(d)
}

func (now *Now) Sunday() time.Time {
	t := now.BeginningOfDay()
	weekday := int(t.Weekday())
	if weekday == 0 {
		return t
	}
	d := time.Duration(7-weekday) * 24 * time.Hour
	return t.Truncate(time.Hour).Add(d)
}

func (now *Now) EndOfSunday() time.Time {
	return now.Sunday().Add(24*time.Hour - time.Nanosecond)
}

func (now *Now) GetCurrentWeeks() (int, int) {
	startOfMonth := now.BeginningOfMonth()

	_, startWeekValue := startOfMonth.ISOWeek()

	_, endWeekValue := now.ISOWeek()
	return startWeekValue, endWeekValue
}

func (now *Now) GetCurrentMonth() (int, int) {
	var (
		startOfYear     = now.BeginningOfYear()
		startMonthValue = startOfYear.Month()
		endMonthValue   = now.Month()
	)
	return int(startMonthValue), int(endMonthValue)
}

func GetStartOfDayUnix() int64 {
	now := time.Now()
	location := time.FixedZone("GMT+7", 7*60*60) // Create a time zone representing GMT+7
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, location)
	return startOfDay.Unix()
}
