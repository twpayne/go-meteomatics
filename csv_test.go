package meteomatics

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientCSVRequest(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/2016-01-20T13:35:00ZP1D:PT3H/t_2m:C,relative_humidity_2m:p/47.423336,9.377225/csv" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(
			"validdate;t_2m:C;relative_humidity_2m:p\n" +
				"2016-01-20T13:35:00Z;-0.829;99.2\n" +
				"2016-01-20T16:35:00Z;-1.574;99.3\n" +
				"2016-01-20T19:35:00Z;-2.167;99\n" +
				"2016-01-20T22:35:00Z;-2.367;98.6\n" +
				"2016-01-21T01:35:00Z;-3.162;95.5\n" +
				"2016-01-21T04:35:00Z;-3.893;75.1\n" +
				"2016-01-21T07:35:00Z;-4.625;79.3\n" +
				"2016-01-21T10:35:00Z;-5.357;100\n" +
				"2016-01-21T13:35:00Z;-6.088;100\n",
		))
	}))

	cr, err := NewClient(
		WithBaseURL(s.URL),
	).CSVRequest(
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
	assert.Equal(t, []ParameterString{"t_2m:C", "relative_humidity_2m:p"}, cr.Parameters)
	assert.Equal(t, 9, len(cr.Rows))
	assert.Equal(t, time.Date(2016, 1, 20, 13, 35, 0, 0, time.UTC), cr.Rows[0].ValidDate)
	assert.Equal(t, []float64{-0.829, 99.2}, cr.Rows[0].Values)
	assert.Equal(t, time.Date(2016, 1, 21, 13, 35, 0, 0, time.UTC), cr.Rows[8].ValidDate)
	assert.Equal(t, []float64{-6.088, 100}, cr.Rows[8].Values)
}

func TestClientCSVRegionRequest(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/2016-12-19T12:00:00Z/t_2m:C/90,-180_-90,180:10x10/csv" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(
			"validdate;2016-12-19 12:00:00\n" +
				"parameter;t_2m:C\n" +
				"data;-180;-140;-100;-60;-20;20;60;100;140;180\n" +
				"90;-15.286;-15.286;-15.286;-15.286;-15.286;-15.286;-15.286;-15.286;-15.286;-15.286\n" +
				"70;-8.161;-16.942;-24.255;-16.692;-0.567;7.3;-22.192;-38.973;-22.411;-8.036\n" +
				"50;2.8;7.1;-6.098;-9.755;9;0.963;-9.348;-15.13;-19.973;2.8\n" +
				"30;20.7;18.8;-5.755;22.1;18.2;16.6;12.5;-3.848;18.2;20.7\n" +
				"10;26.4;26;26.8;27.4;26.7;34.8;26.7;27.8;27.2;26.4\n" +
				"-10;27.1;27;22.5;22.7;24.6;27.4;27.3;27.2;27.9;27.4\n" +
				"-30;21.7;21.9;20.1;22.8;19.7;34.2;22.9;18.7;25.2;21.6\n" +
				"-50;9.7;9.3;8.1;12.2;4.1;1.7;3.7;7.9;8.8;9.7\n" +
				"-70;-1.505;-0.348;-3.13;0.308;-2.036;-2.411;-18.348;-16.411;-25.317;-1.473\n" +
				"-90;-25.63;-25.63;-25.63;-25.63;-25.63;-25.63;-25.63;-25.63;-25.63;-25.63\n",
		))
	}))

	crr, err := NewClient(
		WithBaseURL(s.URL),
	).CSVRegionRequest(
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
	assert.Equal(t, time.Date(2016, 12, 19, 12, 0, 0, 0, time.UTC), crr.ValidDate)
	assert.Equal(t, ParameterString("t_2m:C"), crr.Parameter)
	assert.Equal(t, []float64{-180, -140, -100, -60, -20, 20, 60, 100, 140, 180}, crr.Lons)
	assert.Equal(t, []float64{90, 70, 50, 30, 10, -10, -30, -50, -70, -90}, crr.Lats)
	assert.Equal(t, -15.286, crr.Values[0][0])
	assert.Equal(t, -15.286, crr.Values[0][9])
	assert.Equal(t, -25.63, crr.Values[9][0])
	assert.Equal(t, -25.63, crr.Values[9][9])
}

func TestClientCSVRouteRequest(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/now,now+1H,now+2H/t_2m:C,precip_1h:mm/postal_CH9000+postal_CH8000+postal_CH4000/csv" || r.URL.RawQuery != "route=true" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(
			"lat;lon;validdate;t_2m:C;precip_1h:mm\n" +
				"47.4239;9.3748;2018-10-23T15:47:46Z;10.9;0.02\n" +
				"47.3828;8.5307;2018-10-23T16:47:46Z;12.5;0.00\n" +
				"47.5577;7.5936;2018-10-23T17:47:46Z;13.0;0.00\n",
		))
	}))

	crr, err := NewClient(
		WithBaseURL(s.URL),
	).CSVRouteRequest(
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
	assert.Equal(t, []ParameterString{"t_2m:C", "precip_1h:mm"}, crr.Parameters)
	assert.Equal(t, 3, len(crr.Rows))
	assert.Equal(t, 47.4239, crr.Rows[0].Lat)
	assert.Equal(t, 9.3748, crr.Rows[0].Lon)
	assert.Equal(t, time.Date(2018, 10, 23, 15, 47, 46, 0, time.UTC), crr.Rows[0].ValidDate)
	assert.Equal(t, []float64{10.9, 0.02}, crr.Rows[0].Values)
}
