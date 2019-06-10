# go-meteomatics

[![GoDoc](https://godoc.org/github.com/twpayne/go-meteomatics?status.svg)](https://godoc.org/github.com/twpayne/go-meteomatics)
[![Build Status](https://travis-ci.org/twpayne/go-meteomatics.svg?branch=master)](https://travis-ci.org/twpayne/go-meteomatics)
[![Coverage Status](https://coveralls.io/repos/github/twpayne/go-meteomatics/badge.svg)](https://coveralls.io/github/twpayne/go-meteomatics)

Package `meteomatics` implements a client for the [Meteomatics
API](https://www.meteomatics.com/en/api/overview/).

## Key features

* Idomatic Go API.
* Support for all format types.
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

	jr, err := client.JSONRequest(
		context.Background(),
		meteomatics.TimeNow,
		meteomatics.ParameterString("t_0m:C"),
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

	fmt.Println(jr.Status)
	fmt.Println(jr.Version)
	for _, data := range jr.Data {
		fmt.Println(data.Parameter)
		for _, coordinate := range data.Coordinates {
			fmt.Println(coordinate.StationID)
		}
	}

	// Output:
	// OK
	// 3.0
	// t_0m:C
	// postal_CH9000
}
```

## License

MIT
