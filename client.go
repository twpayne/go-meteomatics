// Package meteomatics is a client library for the Meteomatics API. See
// https://www.meteomatics.com/en/api/overview/.
package meteomatics

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// DefaultBaseURL is the default base URL.
const DefaultBaseURL = "http://api.meteomatics.com"

// A Client is a Client.
type Client struct {
	httpClient      *http.Client
	baseURL         string
	preRequestFuncs []func(*http.Request)
}

// A ClientOption sets an option on a Client.
type ClientOption func(*Client)

// RequestOptions are per-request options.
type RequestOptions struct {
	Source                string
	TemporalInterpolation string
	EnsembleSelect        string
	ClusterSelect         string
	Timeout               int
	Route                 bool
}

// WithBaseURL sets the base URL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// WithHTTPClient sets the http.Client.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithBasicAuth sets the username and password for basic authentication.
func WithBasicAuth(username, password string) ClientOption {
	return func(c *Client) {
		c.preRequestFuncs = append(c.preRequestFuncs, func(req *http.Request) {
			req.SetBasicAuth(username, password)
		})
	}
}

// NewClient returns a new Client with options set.
func NewClient(options ...ClientOption) *Client {
	c := &Client{
		httpClient: http.DefaultClient,
		baseURL:    DefaultBaseURL,
	}
	for _, o := range options {
		o(c)
	}
	return c
}

// RawRequest performs a raw request. It is the caller's responsibility to
// interpret the []byte returned.
func (c *Client) RawRequest(ctx context.Context, time TimeStringer, parameter ParameterStringer, location LocationStringer, format FormatStringer, options *RequestOptions) ([]byte, error) {
	urlStr := fmt.Sprintf("%s/%s/%s/%s/%s", c.baseURL, time.TimeString(), parameter.ParameterString(), location.LocationString(), format.FormatString())
	if values := options.Values(); values != nil {
		urlStr += "?" + values.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Accept", format.ContentType())
	for _, f := range c.preRequestFuncs {
		f(req)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || http.StatusMultipleChoices <= resp.StatusCode {
		respBody, _ := ioutil.ReadAll(resp.Body)
		return nil, &Error{
			Request:      req,
			Response:     resp,
			ResponseBody: respBody,
		}
	}

	return ioutil.ReadAll(resp.Body)
}

// Values returns the url.Values that set the request options defined by o.
func (o *RequestOptions) Values() url.Values {
	if o == nil {
		return nil
	}
	v := url.Values{}
	if o.Source != "" {
		v.Set("source", o.Source)
	}
	if o.TemporalInterpolation != "" && o.TemporalInterpolation != "best" {
		v.Set("temporal_interpolation", o.TemporalInterpolation)
	}
	if o.EnsembleSelect != "" && o.EnsembleSelect != "member:0" {
		v.Set("ens_select", o.EnsembleSelect)
	}
	if o.ClusterSelect != "" {
		v.Set("cluster_select", o.ClusterSelect)
	}
	if o.Timeout != 0 && o.Timeout != 300 {
		// FIXME 30s for WMS/WFS-Queries
		v.Set("timeout", strconv.Itoa(o.Timeout))
	}
	// FIXME route
	if len(v) == 0 {
		return nil
	}
	return v
}
