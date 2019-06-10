package meteomatics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocationString(t *testing.T) {
	for _, tc := range []struct {
		ls       LocationStringer
		expected LocationString
	}{
		{
			ls:       LocationWorld,
			expected: "world",
		},
		{
			ls: Point{
				Lat: 47.419708,
				Lon: 9.358478,
			},
			expected: "47.419708,9.358478",
		},
		{
			ls: PointList{
				Point{
					Lat: 47.41,
					Lon: 9.35,
				},
				Point{
					Lat: 47.51,
					Lon: 8.74,
				},
				Point{
					Lat: 47.13,
					Lon: 8.22,
				},
			},
			expected: "47.41,9.35+47.51,8.74+47.13,8.22",
		},
		{
			ls: Line{
				Start: Point{
					Lat: 50,
					Lon: 10,
				},
				End: Point{
					Lat: 50,
					Lon: 20,
				},
				N: 100,
			},
			expected: "50,10_50,20:100",
		},
		{
			ls: Polyline{
				Start: Point{
					Lat: 50,
					Lon: 10,
				},
				Segments: []PolylineSegment{
					{
						End: Point{
							Lat: 50,
							Lon: 20,
						},
						N: 100,
					},
					{
						End: Point{
							Lat: 60,
							Lon: 20,
						},
						N: 10,
					},
				},
			},
			expected: "50,10_50,20:100+60,20:10",
		},
		{
			ls: Polyline{
				Start: Point{
					Lat: 47.42,
					Lon: 9.37,
				},
				Segments: []PolylineSegment{
					{
						End: Point{
							Lat: 47.46,
							Lon: 9.04,
						},
						N: 10,
					},
					{
						End: Point{
							Lat: 47.51,
							Lon: 8.78,
						},
						N: 10,
					},
					{
						End: Point{
							Lat: 47.39,
							Lon: 8.57,
						},
						N: 10,
					},
				},
			},
			expected: "47.42,9.37_47.46,9.04:10+47.51,8.78:10+47.39,8.57:10",
		},
		{
			ls: RectangleN{
				Min: Point{
					Lat: 40,
					Lon: 10,
				},
				Max: Point{
					Lat: 50,
					Lon: 20,
				},
				NLon: 100,
				NLat: 100,
			},
			expected: "50,10_40,20:100x100",
		},
		{
			ls: RectangleRes{
				Min: Point{
					Lat: 40,
					Lon: 10,
				},
				Max: Point{
					Lat: 50,
					Lon: 20,
				},
				ResLat: 0.1,
				ResLon: 0.1,
			},
			expected: "50,10_40,20:0.1,0.1",
		},
		{
			ls: Postal{
				CountryCode: "CH",
				ZIPCode:     "9014",
			},
			expected: "postal_CH9014",
		},
		{
			ls: Postal{
				CountryCode: "DE",
				ZIPCode:     "10117",
			},
			expected: "postal_DE10117",
		},
		{
			ls: LocationSlice{
				Postal{
					CountryCode: "CH",
					ZIPCode:     "9014",
				},
				Postal{
					CountryCode: "DE",
					ZIPCode:     "10117",
				},
			},
			expected: "postal_CH9014+postal_DE10117",
		},
	} {
		assert.Equal(t, tc.expected, tc.ls.LocationString())
	}
}
