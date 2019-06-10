package meteomatics_test

import (
	"bytes"
	"context"
	"fmt"
	"image/png"
	"os"

	"github.com/twpayne/go-meteomatics"
)

func ExampleClient_CSVRegionRequest() {
	client := meteomatics.NewClient(
		meteomatics.WithBasicAuth(
			os.Getenv("METEOMATICS_USERNAME"),
			os.Getenv("METEOMATICS_PASSWORD"),
		),
	)

	crr, err := client.CSVRegionRequest(
		context.Background(),
		meteomatics.TimeNow,
		meteomatics.ParameterString("t_2m:C"),
		meteomatics.RectangleN{
			Min: meteomatics.Point{
				Lat: -90,
				Lon: -180,
			},
			Max: meteomatics.Point{
				Lat: 90,
				Lon: 180,
			},
			NLat: 10,
			NLon: 10,
		},
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(crr.Parameter)

	// Output:
	// t_2m:C
}

func ExampleClient_JSONRequest() {
	client := meteomatics.NewClient(
		meteomatics.WithBasicAuth(
			os.Getenv("METEOMATICS_USERNAME"),
			os.Getenv("METEOMATICS_PASSWORD"),
		),
	)

	jr, err := client.JSONRequest(
		context.Background(),
		meteomatics.TimeNow,
		meteomatics.ParameterString("t_2m:C"),
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

func ExampleClient_PNGRequest() {
	client := meteomatics.NewClient(
		meteomatics.WithBasicAuth(
			os.Getenv("METEOMATICS_USERNAME"),
			os.Getenv("METEOMATICS_PASSWORD"),
		),
	)

	data, err := client.PNGRequest(
		context.Background(),
		meteomatics.TimeNow,
		meteomatics.ParameterString("t_2m:C"),
		meteomatics.RectangleN{
			Min: meteomatics.Point{
				Lat: -90,
				Lon: -180,
			},
			Max: meteomatics.Point{
				Lat: 90,
				Lon: 180,
			},
			NLon: 10,
			NLat: 10,
		},
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	i, err := png.Decode(bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i.Bounds())

	// Output:
	// (0,0)-(10,10)
}
