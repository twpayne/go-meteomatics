package meteomatics

import (
	"strconv"
	"time"
)

// An IntervalString is a string representation of an interval.
type IntervalString string

// An IntervalStringer can be converted to an IntervalString.
type IntervalStringer interface {
	IntervalString() IntervalString
}

// IntervalString returns s as an IntervalString.
func (s IntervalString) IntervalString() IntervalString {
	return s
}

// Intervals.
const (
	Interval5Min  IntervalString = "5min"
	Interval10Min IntervalString = "10min"
	Interval15Min IntervalString = "15min"
	Interval30Min IntervalString = "30min"
	Interval1H    IntervalString = "1h"
	Interval3H    IntervalString = "3h"
	Interval6H    IntervalString = "6h"
	Interval12H   IntervalString = "12h"
	Interval24H   IntervalString = "24h"
)

// An Interval is a time.Duration.
type Interval time.Duration

// IntervalString returns i as an IntervalString.
func (i Interval) IntervalString() IntervalString {
	if time.Duration(i)%time.Hour == 0 {
		return IntervalString(strconv.Itoa(int(time.Duration(i)/time.Hour)) + "h")
	}
	return IntervalString(strconv.Itoa(int(time.Duration(i)/time.Minute)) + "min")
}
