package meteomatics

import (
	"context"
	"encoding/json"
	"time"
)

// A JSONDate is a value at a date.
type JSONDate struct {
	Date  time.Time `json:"date"`
	Value float64   `json:"value"`
}

// A JSONCoordinates is a series of values.
type JSONCoordinates struct {
	Dates     []JSONDate `json:"dates"`
	Lat       float64    `json:"lat"`
	Lon       float64    `json:"lon"`
	StationID string     `json:"station_id"`
}

// A JSONData is a parameter measured at a coordinate.
type JSONData struct {
	Coordinates []JSONCoordinates `json:"coordinates"`
	Parameter   ParameterString   `json:"parameter"`
}

// A JSONResponse is a JSON response.
type JSONResponse struct {
	Version       string     `json:"version"`
	User          string     `json:"user"`
	DateGenerated time.Time  `json:"dateGenerated"`
	Status        string     `json:"status"`
	Data          []JSONData `json:"data"`
}

// A JSONRouteParameter is a JSON route parameter.
type JSONRouteParameter struct {
	Parameter ParameterString `json:"parameter"`
	Value     float64         `json:"value"`
}

// A JSONRouteData is a JSON route data.
type JSONRouteData struct {
	Lat        float64              `json:"lat"`
	Lon        float64              `json:"lon"`
	Date       time.Time            `json:"date"`
	Parameters []JSONRouteParameter `json:"parameters"`
}

// A JSONRouteResponse is a response to a JSON route request.
type JSONRouteResponse struct {
	Version       string          `json:"version"`
	User          string          `json:"user"`
	DateGenerated time.Time       `json:"dateGenerated"`
	Status        string          `json:"status"`
	Data          []JSONRouteData `json:"data"`
}

// RequestJSON requests a forecast in JSON format.
func (c *Client) RequestJSON(ctx context.Context, ts TimeStringer, ps ParameterStringer, ls LocationStringer, options *RequestOptions) (*JSONResponse, error) {
	data, err := c.Request(ctx, ts, ps, ls, FormatJSON, options)
	if err != nil {
		return nil, err
	}
	jr := &JSONResponse{}
	if err := json.Unmarshal(data, jr); err != nil {
		return nil, err
	}
	if jr.Status != "OK" {
		return nil, jr
	}
	return jr, nil
}

// RequestJSONRoute requests a forecast in JSON format.
func (c *Client) RequestJSONRoute(ctx context.Context, ts TimeStringer, ps ParameterStringer, ls LocationStringer, options *RequestOptions) (*JSONRouteResponse, error) {
	var ro RequestOptions
	if options != nil {
		ro = *options
	}
	ro.Route = true
	data, err := c.Request(ctx, ts, ps, ls, FormatJSON, &ro)
	if err != nil {
		return nil, err
	}
	jrr := &JSONRouteResponse{}
	if err := json.Unmarshal(data, jrr); err != nil {
		return nil, err
	}
	if jrr.Status != "OK" {
		return nil, jrr
	}
	return jrr, nil
}

func (r *JSONResponse) Error() string {
	return r.Status
}

func (r *JSONRouteResponse) Error() string {
	return r.Status
}
