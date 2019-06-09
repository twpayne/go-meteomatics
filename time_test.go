package meteomatics

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeString(t *testing.T) {
	for _, tc := range []struct {
		ts       TimeStringer
		expected TimeString
	}{
		{
			ts:       TimeNow,
			expected: "now",
		},
		{
			ts:       TimeTomorrow,
			expected: "tomorrow",
		},
		{
			ts:       TimeYesterday,
			expected: "yesterday",
		},
		{
			ts: TimePeriod{
				Start:    time.Date(2017, 5, 28, 13, 0, 0, 0, time.UTC),
				Duration: 10 * 24 * time.Hour,
				Step:     time.Hour,
			},
			expected: "2017-05-28T13:00:00ZP10D:PT1H",
		},
		{
			ts: TimePoint{
				Time: time.Date(2015, 1, 20, 18, 0, 0, 0, time.UTC),
			},
			expected: "2015-01-20T18:00:00Z",
		},
		{
			ts: TimeRange{
				Start: time.Date(2017, 5, 28, 13, 0, 0, 0, time.UTC),
				End:   time.Date(2017, 5, 30, 13, 0, 0, 0, time.UTC),
				Step:  24 * time.Hour,
			},
			expected: "2017-05-28T13:00:00Z--2017-05-30T13:00:00Z:P1D",
		},
		{
			ts: TimeSlice{
				TimePoint{
					Time: time.Date(2018, 10, 20, 18, 0, 0, 0, time.UTC),
				},
				TimePoint{
					Time: time.Date(2018, 10, 21, 18, 0, 0, 0, time.UTC),
				},
				TimePoint{
					Time: time.Date(2018, 10, 22, 18, 0, 0, 0, time.UTC),
				},
			},
			expected: "2018-10-20T18:00:00Z,2018-10-21T18:00:00Z,2018-10-22T18:00:00Z",
		},
		{
			ts: TimeSlice{
				TimePoint{
					Time: time.Date(2018, 10, 20, 18, 0, 0, 0, time.UTC),
				},
				TimePeriod{
					Start:    time.Date(2018, 10, 21, 18, 0, 0, 0, time.UTC),
					Duration: 2 * time.Hour,
					Step:     20 * time.Minute,
				},
			},
			expected: "2018-10-20T18:00:00Z,2018-10-21T18:00:00ZPT2H:PT20M",
		},
		{
			ts: TimeSlice{
				TimePeriod{
					Start:    time.Date(2018, 10, 20, 18, 0, 0, 0, time.UTC),
					Duration: 2 * time.Hour,
					Step:     20 * time.Minute,
				},
				TimePeriod{
					Start:    time.Date(2018, 10, 21, 18, 0, 0, 0, time.UTC),
					Duration: 2 * time.Hour,
					Step:     20 * time.Minute,
				},
			},
			expected: "2018-10-20T18:00:00ZPT2H:PT20M,2018-10-21T18:00:00ZPT2H:PT20M",
		},
	} {
		assert.Equal(t, tc.expected, tc.ts.TimeString())
	}
}
