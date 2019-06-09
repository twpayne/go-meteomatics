package meteomatics_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/twpayne/go-meteomatics"
)

func mustNewClient(t *testing.T) *meteomatics.Client {
	return meteomatics.NewClient(
		meteomatics.WithBasicAuth(
			os.Getenv("METEOMATICS_USERNAME"),
			os.Getenv("METEOMATICS_PASSWORD"),
		),
	)
}

func TestClientSimple(t *testing.T) {
	ctx := context.Background()
	c := mustNewClient(t)
	jr, err := c.JSONRequest(
		ctx,
		meteomatics.TimeNow,
		meteomatics.ParameterString("t_0m:C"),
		meteomatics.Postal{
			CountryCode: "CH",
			ZIPCode:     "8000",
		},
		nil,
	)
	require.NoError(t, err)
	assert.Equal(t, "OK", jr.Status)
	assert.Equal(t, "3.0", jr.Version)
	require.Equal(t, len(jr.Datas), 1)
	assert.Equal(t, "t_0m:C", jr.Datas[0].Parameter)
	require.Equal(t, 1, len(jr.Datas[0].Coordinates))
	assert.Equal(t, "postal_CH8000", jr.Datas[0].Coordinates[0].StationID)
}
