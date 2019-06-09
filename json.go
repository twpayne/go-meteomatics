package meteomatics

import (
	"context"
	"encoding/json"
	"time"
)

// A JSONCoordinateValues is a series of values.
type JSONCoordinateValues struct {
	StationID string           `json:"station_id"`
	Dates     []*JSONDateValue `json:"dates"`
}

// A JSONData is a parameter measured at a coordinate.
type JSONData struct {
	Parameter   string                  `json:"parameter"`
	Coordinates []*JSONCoordinateValues `json:"coordinates"`
}

// A JSONDateValue is a value at a date.
type JSONDateValue struct {
	Date  *time.Time `json:"date"`
	Value float64    `json:"value"`
}

// A JSONResponse is a JSON response.
type JSONResponse struct {
	Version       string      `json:"version"`
	User          string      `json:"user"`
	DateGenerated *time.Time  `json:"dateGenerated"`
	Status        string      `json:"status"`
	Datas         []*JSONData `json:"data"`
}

// JSONRequest requests a forecast in JSON format.
func (c *Client) JSONRequest(ctx context.Context, time TimeStringer, parameter ParameterStringer, location LocationStringer, options *RequestOptions) (*JSONResponse, error) {
	data, err := c.RawRequest(ctx, time, parameter, location, FormatJSON, options)
	if err != nil {
		return nil, err
	}
	var jsonResponse JSONResponse
	if err := json.Unmarshal(data, &jsonResponse); err != nil {
		return nil, err
	}
	return &jsonResponse, nil
}
