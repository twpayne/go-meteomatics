# go-meteomatics

[![GoDoc](https://godoc.org/github.com/twpayne/go-meteomatics?status.svg)](https://godoc.org/github.com/twpayne/go-meteomatics)
[![Build Status](https://travis-ci.org/twpayne/go-meteomatics.svg?branch=master)](https://travis-ci.org/twpayne/go-meteomatics)
[![Coverage Status](https://coveralls.io/repos/github/twpayne/go-meteomatics/badge.svg)](https://coveralls.io/github/twpayne/go-meteomatics)

Package `meteomatics` implements a client for the [Meteomatics
API](https://www.meteomatics.com/en/api/overview/).

## Key features

* Idomatic Go API.
* Support for CSV, JSON, and PNG formats.
* Support for all location types.
* Support for all parameters.
* Support for all time types.
* Support for `context`.
* Support for Go modules.

## Example

```go
func ExampleNewClient_simple() {
	client := meteomatics.NewClient(
		meteomatics.WithBasicAuth(
			os.Getenv("METEOMATICS_USERNAME"),
			os.Getenv("METEOMATICS_PASSWORD"),
		),
	)

	cr, err := client.CSVRequest(
		context.Background(),
		meteomatics.TimeSlice{
			meteomatics.TimeNow,
			meteomatics.NowOffset(1 * time.Hour),
		},
		meteomatics.Parameter{
			Name:  meteomatics.ParameterTemperature,
			Level: meteomatics.LevelMeters(2),
			Units: meteomatics.UnitsCelsius,
		},
		meteomatics.Postal{
			CountryCode: "CH",
			ZIPCode:     "9000",
		},
		&meteomatics.RequestOptions{},
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cr.Parameters)
	for _, row := range cr.Rows {
		fmt.Println(row.ValidDate)
		fmt.Println(row.Values)
	}
}
```

## License

MIT
