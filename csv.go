package meteomatics

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"strconv"
	"strings"
	"time"
)

var errCSVParse = errors.New("csv parse error")

// A CSVRow is a CSV row.
type CSVRow struct {
	ValidDate time.Time
	Values    []float64
}

// A CSVResponse is a response to a CSV request.
type CSVResponse struct {
	Header []ParameterString
	Rows   []CSVRow
}

// A CSVRegionResponse is a response to a CSV region request.
type CSVRegionResponse struct {
	ValidDate time.Time
	Parameter ParameterString
	Lats      []float64
	Lons      []float64
	Values    [][]float64
}

// CSVRequest requests a forecast in CSV format.
func (c *Client) CSVRequest(ctx context.Context, ts TimeStringer, ps ParameterStringer, ls LocationStringer, options *RequestOptions) (*CSVResponse, error) {
	data, err := c.RawRequest(ctx, ts, ps, ls, FormatCSV, options)
	if err != nil {
		return nil, err
	}

	cr := &CSVResponse{}

	s := bufio.NewScanner(bytes.NewReader(data))
	if !s.Scan() {
		return nil, errCSVParse
	}
	record := strings.Split(s.Text(), ";")
	if len(record) < 1 {
		return nil, errCSVParse
	}
	if record[0] != "validdate" {
		return nil, errCSVParse
	}
	cols := len(record)
	cr.Header = make([]ParameterString, 0, cols-1)
	for i := 1; i < cols; i++ {
		cr.Header = append(cr.Header, ParameterString(record[i]))
	}

	for s.Scan() {
		record := strings.Split(s.Text(), ";")
		if len(record) != cols {
			return nil, errCSVParse
		}
		var row CSVRow
		row.ValidDate, err = time.Parse(time.RFC3339, record[0])
		if err != nil {
			return nil, err
		}
		row.Values = make([]float64, 0, cols-1)
		for i := 1; i < cols; i++ {
			value, err := strconv.ParseFloat(record[i], 64)
			if err != nil {
				return nil, err
			}
			row.Values = append(row.Values, value)
		}
		cr.Rows = append(cr.Rows, row)
	}

	return cr, s.Err()
}

// CSVRegionRequest requests a region forecast in CSV format.
func (c *Client) CSVRegionRequest(ctx context.Context, ts TimeStringer, ps ParameterStringer, ls LocationStringer, options *RequestOptions) (*CSVRegionResponse, error) {
	data, err := c.RawRequest(ctx, ts, ps, ls, FormatCSV, options)
	if err != nil {
		return nil, err
	}

	crr := &CSVRegionResponse{}

	s := bufio.NewScanner(bytes.NewReader(data))

	validDate, err := scanRow(s, "validdate")
	if err != nil {
		return nil, err
	}
	crr.ValidDate, err = time.Parse("2006-01-02 15:04:05", validDate)
	if err != nil {
		return nil, err
	}

	parameter, err := scanRow(s, "parameter")
	if err != nil {
		return nil, err
	}
	crr.Parameter = ParameterString(parameter)

	if !s.Scan() {
		return nil, errCSVParse
	}
	record := strings.Split(s.Text(), ";")
	if len(record) == 0 || record[0] != "data" {
		return nil, errCSVParse
	}
	cols := len(record)
	crr.Lons = make([]float64, 0, cols-1)
	for i := 1; i < cols; i++ {
		lon, err := strconv.ParseFloat(record[i], 64)
		if err != nil {
			return nil, err
		}
		crr.Lons = append(crr.Lons, lon)
	}

	for s.Scan() {
		record := strings.Split(s.Text(), ";")
		if len(record) != cols {
			return nil, errCSVParse
		}
		lat, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			return nil, err
		}
		crr.Lats = append(crr.Lats, lat)
		values := make([]float64, 0, cols-1)
		for i := 1; i < cols; i++ {
			value, err := strconv.ParseFloat(record[i], 64)
			if err != nil {
				return nil, err
			}
			values = append(values, value)
		}
		crr.Values = append(crr.Values, values)
	}

	return crr, s.Err()
}

func scanRow(s *bufio.Scanner, name string) (string, error) {
	if !s.Scan() {
		return "", errCSVParse
	}
	record := strings.Split(s.Text(), ";")
	if len(record) != 2 || record[0] != name {
		return "", errCSVParse
	}
	return record[1], nil
}
