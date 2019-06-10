package meteomatics

import (
	"fmt"
	"net/http"
)

// An Error is an error.
type Error struct {
	Request      *http.Request
	Response     *http.Response
	ResponseBody []byte
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %d %s", e.Request.URL, e.Response.StatusCode, http.StatusText(e.Response.StatusCode))
}
