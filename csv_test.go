package meteomatics

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientRequestCSV(t *testing.T) {
	s := newTestServer(
		t,
		"/2016-01-20T13:35:00ZP1D:PT3H/t_2m:C,relative_humidity_2m:p/47.423336,9.377225/csv",
		"testdata/temperature_and_relative_humidity_time_series.csv",
	)
	r, err := NewClient(WithBaseURL(s.URL)).RequestCSV(
		context.Background(),
		TimePeriod{
			Start:    time.Date(2016, 1, 20, 13, 35, 0, 0, time.UTC),
			Duration: 24 * time.Hour,
			Step:     3 * time.Hour,
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
		Point{
			Lat: 47.423336,
			Lon: 9.377225,
		},
		nil,
	)
	require.NoError(t, err)
	assert.Equal(t, []ParameterString{"t_2m:C", "relative_humidity_2m:p"}, r.Parameters)
	require.Len(t, r.Rows, 9)
	assert.Equal(t, time.Date(2016, 1, 20, 13, 35, 0, 0, time.UTC), r.Rows[0].ValidDate)
	assert.Equal(t, []float64{-0.829, 99.2}, r.Rows[0].Values)
	assert.Equal(t, time.Date(2016, 1, 21, 13, 35, 0, 0, time.UTC), r.Rows[8].ValidDate)
	assert.Equal(t, []float64{-6.088, 100}, r.Rows[8].Values)
}

func TestClientRequestCSVRegion(t *testing.T) {
	s := newTestServer(
		t,
		"/2016-12-19T12:00:00Z/t_2m:C/90,-180_-90,180:10x10/csv",
		"testdata/temperature_geographical_region.csv",
	)
	r, err := NewClient(WithBaseURL(s.URL)).RequestCSVRegion(
		context.Background(),
		Time(time.Date(2016, 12, 19, 12, 00, 0, 0, time.UTC)),
		Parameter{
			Name:  ParameterTemperature,
			Level: LevelMeters(2),
			Units: UnitsCelsius,
		},
		RectangleN{
			Min: Point{
				Lat: -90,
				Lon: -180,
			},
			Max: Point{
				Lat: 90,
				Lon: 180,
			},
			NLat: 10,
			NLon: 10,
		},
		nil,
	)
	require.NoError(t, err)
	assert.Equal(t, time.Date(2016, 12, 19, 12, 0, 0, 0, time.UTC), r.ValidDate)
	assert.Equal(t, ParameterString("t_2m:C"), r.Parameter)
	assert.Equal(t, []float64{-180, -140, -100, -60, -20, 20, 60, 100, 140, 180}, r.Lons)
	assert.Equal(t, []float64{90, 70, 50, 30, 10, -10, -30, -50, -70, -90}, r.Lats)
	assert.Equal(t, -15.286, r.Values[0][0])
	assert.Equal(t, -15.286, r.Values[0][9])
	assert.Equal(t, -25.63, r.Values[9][0])
	assert.Equal(t, -25.63, r.Values[9][9])
}

func TestClientRequestCSVRoute(t *testing.T) {
	s := newTestServer(
		t,
		"/now,now+1H,now+2H/t_2m:C,precip_1h:mm/postal_CH9000+postal_CH8000+postal_CH4000/csv?route=true",
		"testdata/csv_route_query.csv",
	)
	r, err := NewClient(WithBaseURL(s.URL)).RequestCSVRoute(
		context.Background(),
		TimeSlice{
			TimeNow,
			NowOffset(1 * time.Hour),
			NowOffset(2 * time.Hour),
		},
		ParameterSlice{
			Parameter{
				Name:  ParameterTemperature,
				Level: LevelMeters(2),
				Units: UnitsCelsius,
			},
			Parameter{
				Name:     ParameterPrecipitation,
				Interval: Interval(1 * time.Hour),
				Units:    UnitsMillimeters,
			},
		},
		LocationSlice{
			Postal{
				CountryCode: "CH",
				ZIPCode:     "9000",
			},
			Postal{
				CountryCode: "CH",
				ZIPCode:     "8000",
			},
			Postal{
				CountryCode: "CH",
				ZIPCode:     "4000",
			},
		},
		nil,
	)
	require.NoError(t, err)
	assert.Equal(t, []ParameterString{"t_2m:C", "precip_1h:mm"}, r.Parameters)
	require.Len(t, r.Rows, 3)
	assert.Equal(t, 47.4239, r.Rows[0].Lat)
	assert.Equal(t, 9.3748, r.Rows[0].Lon)
	assert.Equal(t, time.Date(2018, 10, 23, 15, 47, 46, 0, time.UTC), r.Rows[0].ValidDate)
	assert.Equal(t, []float64{10.9, 0.02}, r.Rows[0].Values)
}
