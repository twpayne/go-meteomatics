package meteomatics

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientRequestJSON(t *testing.T) {
	s := newTestServer(
		t,
		"/2016-12-20T00:00:00ZP2D:P1D/t_2m:C,relative_humidity_2m:p/50,10+40,20/json",
		"testdata/temperature_and_relative_humidity_between_two_times_at_two_locations.json",
	)
	r, err := NewClient(WithBaseURL(s.URL)).RequestJSON(
		context.Background(),
		TimePeriod{
			Start:    time.Date(2016, 12, 20, 0, 0, 0, 0, time.UTC),
			Duration: 2 * 24 * time.Hour,
			Step:     24 * time.Hour,
		},
		ParameterSlice{
			Parameter{
				Name:  ParameterTemperature,
				Level: LevelMeters(2),
				Units: UnitsCelsius,
			},
			Parameter{
				Name:  ParameterRelativeHumidity,
				Level: LevelMeters(2),
				Units: UnitsPercentage,
			},
		},
		LocationSlice{
			Point{
				Lat: 50,
				Lon: 10,
			},
			Point{
				Lat: 40,
				Lon: 20,
			},
		},
		&RequestOptions{},
	)
	require.NoError(t, err)
	assert.Equal(t, "3.0", r.Version)
	assert.Equal(t, "internal-api-beta-user", r.User)
	assert.Equal(t, time.Date(2016, 12, 23, 15, 24, 7, 0, time.UTC), r.DateGenerated)
	assert.Equal(t, "OK", r.Status)
	require.Len(t, r.Data, 2)
	assert.Equal(t, ParameterString("t_2m:C"), r.Data[0].Parameter)
	require.Equal(t, 2, len(r.Data[0].Coordinates))
	assert.Equal(t, 50.0, r.Data[0].Coordinates[0].Lat)
	assert.Equal(t, 10.0, r.Data[0].Coordinates[0].Lon)
	require.Len(t, r.Data[0].Coordinates[0].Dates, 3)
	assert.Equal(t, time.Date(2016, 12, 20, 0, 0, 0, 0, time.UTC), r.Data[0].Coordinates[0].Dates[0].Date)
	assert.Equal(t, -1.18699, r.Data[0].Coordinates[0].Dates[0].Value)
	assert.Equal(t, ParameterString("relative_humidity_2m:p"), r.Data[1].Parameter)
	require.Len(t, r.Data[1].Coordinates, 2)
	assert.Equal(t, 40.0, r.Data[1].Coordinates[1].Lat)
	assert.Equal(t, 20.0, r.Data[1].Coordinates[1].Lon)
	require.Len(t, r.Data[1].Coordinates[1].Dates, 3)
	assert.Equal(t, time.Date(2016, 12, 22, 0, 0, 0, 0, time.UTC), r.Data[1].Coordinates[1].Dates[2].Date)
	assert.Equal(t, 64.9726, r.Data[1].Coordinates[1].Dates[2].Value)
}

func TestClientRequestJSONRoute(t *testing.T) {
	s := newTestServer(
		t,
		"/2018-10-19T12:00:00ZPT1H:PT30M/t_2m:C,wind_speed_10m:ms/47,9_45,7:3/json?route=true",
		"testdata/json_route_query.json",
	)
	r, err := NewClient(WithBaseURL(s.URL)).RequestJSONRoute(
		context.Background(),
		TimePeriod{
			Start:    time.Date(2018, 10, 19, 12, 0, 0, 0, time.UTC),
			Duration: 1 * time.Hour,
			Step:     30 * time.Minute,
		},
		ParameterSlice{
			Parameter{
				Name:  ParameterTemperature,
				Level: LevelMeters(2),
				Units: UnitsCelsius,
			},
			Parameter{
				Name:  ParameterWindSpeed,
				Level: LevelMeters(10),
				Units: UnitsMetersPerSecond,
			},
		},
		Polyline{
			Start: Point{
				Lat: 47,
				Lon: 9,
			},
			Segments: []PolylineSegment{
				{
					End: Point{
						Lat: 45,
						Lon: 7,
					},
					N: 3,
				},
			},
		},
		&RequestOptions{},
	)
	require.NoError(t, err)
	assert.Equal(t, "3.0", r.Version)
	assert.Equal(t, "internal-api-beta-user", r.User)
	assert.Equal(t, time.Date(2018, 10, 19, 9, 13, 4, 0, time.UTC), r.DateGenerated)
	assert.Equal(t, "OK", r.Status)
	assert.Len(t, r.Data, 3)
	assert.Equal(t, 47.0, r.Data[0].Lat)
	assert.Equal(t, 9.0, r.Data[0].Lon)
	assert.Equal(t, time.Date(2018, 10, 19, 12, 0, 0, 0, time.UTC), r.Data[0].Date)
	assert.Len(t, r.Data[0].Parameters, 2)
	assert.Equal(t, ParameterString("t_2m:C"), r.Data[0].Parameters[0].Parameter)
	assert.Equal(t, 3.9, r.Data[0].Parameters[0].Value)
	assert.Equal(t, 45.0, r.Data[2].Lat)
	assert.Equal(t, 7.0, r.Data[2].Lon)
	assert.Equal(t, time.Date(2018, 10, 19, 13, 0, 0, 0, time.UTC), r.Data[2].Date)
	assert.Len(t, r.Data[2].Parameters, 2)
	assert.Equal(t, ParameterString("wind_speed_10m:ms"), r.Data[2].Parameters[1].Parameter)
	assert.Equal(t, 2.0, r.Data[2].Parameters[1].Value)
}

func TestClientRequestJSONError(t *testing.T) {
	s := newTestServer(t, "/now/t_2m:C/0,190/json", "testdata/out_of_range_error.json")
	_, err := NewClient(WithBaseURL(s.URL)).RequestJSON(
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
		&RequestOptions{},
	)
	require.Error(t, err)
	assert.True(t, strings.HasPrefix(err.Error(), "Not enough data outside temporal and/or spatial domain"))
}
