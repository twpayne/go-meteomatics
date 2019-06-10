package meteomatics

import (
	"context"
	"encoding/json"
	"time"
)

// A JSONDate is a value at a date.
type JSONDate struct {
	Date  *time.Time `json:"date"`
	Value float64    `json:"value"`
}

// A JSONCoordinates is a series of values.
type JSONCoordinates struct {
	Dates     []*JSONDate `json:"dates"`
	Lat       float64     `json:"lat"`
	Lon       float64     `json:"lon"`
	StationID string      `json:"station_id"`
}

// A JSONData is a parameter measured at a coordinate.
type JSONData struct {
	Coordinates []*JSONCoordinates `json:"coordinates"`
	Parameter   string             `json:"parameter"`
}

// A JSONResponse is a JSON response.
type JSONResponse struct {
	Data          []*JSONData `json:"data"`
	DateGenerated *time.Time  `json:"dateGenerated"`
	Status        string      `json:"status"`
	User          string      `json:"user"`
	Version       string      `json:"version"`
}

// JSONRequest requests a forecast in JSON format.
func (c *Client) JSONRequest(ctx context.Context, ts TimeStringer, ps ParameterStringer, ls LocationStringer, options *RequestOptions) (*JSONResponse, error) {
	data, err := c.Request(ctx, ts, ps, ls, FormatJSON, options)
	if err != nil {
		return nil, err
	}
	var jsonResponse JSONResponse
	if err := json.Unmarshal(data, &jsonResponse); err != nil {
		return nil, err
	}
	return &jsonResponse, nil
}
