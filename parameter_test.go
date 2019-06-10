package meteomatics

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParameterString(t *testing.T) {
	for _, tc := range []struct {
		ps       ParameterStringer
		expected ParameterString
	}{
		{
			ps:       ParameterString("t_2m:C"),
			expected: "t_2m:C",
		},
		{
			ps: Parameter{
				Name:  ParameterTemperature,
				Level: LevelMeters(2),
				Units: UnitsCelsius,
			},
			expected: "t_2m:C",
		},
		{
			ps: Parameter{
				Name:     ParameterTemperatureMean,
				Level:    LevelHectopascals(500),
				Interval: Interval(6 * time.Hour),
				Units:    UnitsKelvin,
			},
			expected: "t_mean_500hPa_6h:K",
		},
		{
			ps: Parameter{
				Name:  ParameterRelativeHumidity,
				Level: LevelHectopascals(1000),
				Units: UnitsPercentage,
			},
			expected: "relative_humidity_1000hPa:p",
		},
		{
			ps: Parameter{
				Name:     ParameterPrecipitation,
				Interval: Interval1H,
				Units:    UnitsMillimeters,
			},
			expected: "precip_1h:mm",
		},
		{
			ps: Parameter{
				Name:  ParameterPrecipitationType,
				Units: UnitsIndex,
			},
			expected: "precip_type:idx",
		},
	} {
		assert.Equal(t, tc.expected, tc.ps.ParameterString())
	}
}
