package meteomatics

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientJSONRequest(t *testing.T) {
	s := newTestServer(t, "/now/t_2m:C/postal_CH9000/json", "testdata/example.json")
	jr, err := NewClient(WithBaseURL(s.URL)).JSONRequest(
		context.Background(),
		TimeNow,
		Parameter{
			Name:  ParameterTemperature,
			Level: LevelMeters(2),
			Units: UnitsCelsius,
		},
		Postal{
			CountryCode: "CH",
			ZIPCode:     "9000",
		},
		nil,
	)
	require.NoError(t, err)
	assert.Equal(t, "3.0", jr.Version)
	assert.Equal(t, "username", jr.User)
	assert.Equal(t, time.Date(2019, 6, 10, 16, 3, 40, 0, time.UTC), jr.DateGenerated)
	assert.Equal(t, "OK", jr.Status)
	require.Equal(t, 1, len(jr.Data))
	assert.Equal(t, "t_2m:C", jr.Data[0].Parameter)
	require.Equal(t, 1, len(jr.Data[0].Coordinates))
	assert.Equal(t, "postal_CH9000", jr.Data[0].Coordinates[0].StationID)
	require.Equal(t, 1, len(jr.Data[0].Coordinates[0].Dates))
	assert.Equal(t, time.Date(2019, 6, 10, 16, 3, 40, 0, time.UTC), jr.Data[0].Coordinates[0].Dates[0].Date)
	assert.Equal(t, 15.6, jr.Data[0].Coordinates[0].Dates[0].Value)
}

func TestClientJSONRequestError(t *testing.T) {
	s := newTestServer(t, "/now/t_2m:C/0,190/json", "testdata/out_of_range_error.json")
	_, err := NewClient(WithBaseURL(s.URL)).JSONRequest(
		context.Background(),
		TimeNow,
		Parameter{
			Name:  ParameterTemperature,
			Level: LevelMeters(2),
			Units: UnitsCelsius,
		},
		Point{
			Lat: 0,
			Lon: 190,
		},
		nil,
	)
	require.Error(t, err)
	assert.True(t, strings.HasPrefix(err.Error(), "Not enough data outside temporal and/or spatial domain"))
}
