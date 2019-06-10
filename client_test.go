package meteomatics

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientRequestError(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))

	_, err := NewClient(WithBaseURL(s.URL)).Request(
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
		FormatCSV,
		nil,
	)
	require.Error(t, err)
	assert.Equal(t, s.URL+"/now/t_2m:C/0,190/csv: 404 Not Found", err.Error())
}
