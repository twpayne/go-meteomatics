package meteomatics

import (
	"strconv"
	"strings"
	"time"
)

// A TimeString is a string representing a time.
type TimeString string

// A TimeStringer can be converted to a TimeString.
type TimeStringer interface {
	TimeString() TimeString
}

// Time shortcuts.
const (
	TimeNow       TimeString = "now"
	TimeTomorrow  TimeString = "tomorrow"
	TimeYesterday TimeString = "yesterday"
)

// TimeString returns s as a TimeString.
func (s TimeString) TimeString() TimeString {
	return s
}

type TimePeriod struct {
	Start    time.Time
	Duration time.Duration
	Step     time.Duration
}

func (p TimePeriod) TimeString() TimeString {
	return TimeString(formatTime(p.Start) +
		"P" + formatDuration(p.Duration) +
		":P" + formatDuration(p.Step))
}

type TimePoint struct {
	time.Time
}

func (p TimePoint) TimeString() TimeString {
	return TimeString(formatTime(p.Time))
}

type TimeRange struct {
	Start time.Time
	End   time.Time
	Step  time.Duration
}

func (r TimeRange) TimeString() TimeString {
	return TimeString(formatTime(r.Start) +
		"--" + formatTime(r.End) +
		":P" + formatDuration(r.Step))
}

type TimeSlice []TimeStringer

func (s TimeSlice) TimeString() TimeString {
	ss := make([]string, len(s))
	for i, ts := range s {
		ss[i] = string(ts.TimeString())
	}
	return TimeString(strings.Join(ss, ","))
}

func formatDuration(d time.Duration) string {
	for _, unit := range []struct {
		divisor time.Duration
		prefix  string
		suffix  string
	}{
		{divisor: 7 * 24 * time.Hour, suffix: "W"},
		{divisor: 24 * time.Hour, suffix: "D"},
		{divisor: time.Hour, prefix: "T", suffix: "H"},
		{divisor: time.Minute, prefix: "T", suffix: "M"},
	} {
		if d%unit.divisor == 0 {
			return unit.prefix + strconv.Itoa(int(d/unit.divisor)) + unit.suffix
		}
	}
	return "T" + strconv.Itoa(int(d/time.Second)) + "S"
}

func formatTime(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}
