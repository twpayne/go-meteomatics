package meteomatics

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, expectedPath, bodyFilename string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if r.URL.RawQuery != "" {
			path += "?" + r.URL.RawQuery
		}
		require.Equal(t, expectedPath, path)
		w.WriteHeader(http.StatusOK)
		body, err := ioutil.ReadFile(bodyFilename)
		require.NoError(t, err)
		n, err := w.Write(body)
		require.NoError(t, err)
		require.Equal(t, len(body), n)
	}))
}
