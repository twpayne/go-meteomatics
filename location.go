package meteomatics

import (
	"strconv"
	"strings"
)

// A LocationString is a string representing a location.
type LocationString string

// A LocationStringer can be converted to a LocationString.
type LocationStringer interface {
	LocationString() LocationString
}

// Location shortcuts.
const (
	LocationWorld        LocationString = "world"
	LocationGlobal       LocationString = "global"
	LocationAfrica       LocationString = "africa"
	LocationAsia         LocationString = "asia"
	LocationAustralia    LocationString = "australia"
	LocationEurope       LocationString = "europe"
	LocationNorthAmerica LocationString = "north-america"
	LocationSouthAmerica LocationString = "south-america"
)

// LocationString returns s as a LocationString.
func (s LocationString) LocationString() LocationString {
	return s
}

// A Point is a point.
type Point struct {
	Lat float64
	Lon float64
}

// LocationString returns p as a LocationString.
func (p Point) LocationString() LocationString {
	return formatFloat(p.Lat) + "," + formatFloat(p.Lon)
}

// A PointList is a list of Points.
type PointList []Point

// LocationString returns l as a LocationString.
func (l PointList) LocationString() LocationString {
	ss := make([]string, len(l))
	for i, p := range l {
		ss[i] = string(p.LocationString())
	}
	return LocationString(strings.Join(ss, "+"))
}

// A Line is a line.
type Line struct {
	Start Point
	End   Point
	N     int
}

// LocationString returns l as a LocationString.
func (l Line) LocationString() LocationString {
	return l.Start.LocationString() +
		"_" + l.End.LocationString() +
		":" + formatInt(l.N)
}

// A PolylineSegment is a segment of a Polyline.
type PolylineSegment struct {
	End Point
	N   int
}

// A Polyline is a polyline.
type Polyline struct {
	Start    Point
	Segments []PolylineSegment
}

// LocationString returns p as a LocationString.
func (p Polyline) LocationString() LocationString {
	b := strings.Builder{}
	_, _ = b.WriteString(string(Line{
		Start: p.Start,
		End:   p.Segments[0].End,
		N:     p.Segments[0].N,
	}.LocationString()))
	for i := 1; i < len(p.Segments); i++ {
		s := p.Segments[i]
		_, _ = b.WriteString(string("+" + s.End.LocationString() + ":" + formatInt(s.N)))
	}
	return LocationString(b.String())
}

// A RectangleN is a rectangle with a number of points.
type RectangleN struct {
	Min  Point
	Max  Point
	NLon int
	NLat int
}

// LocationString returns r as a LocationString.
func (r RectangleN) LocationString() LocationString {
	return formatFloat(r.Max.Lat) + "," + formatFloat(r.Min.Lon) +
		"_" + formatFloat(r.Min.Lat) + "," + formatFloat(r.Max.Lon) +
		":" + formatInt(r.NLon) + "x" + formatInt(r.NLat)
}

// A RectangleRes is a rectangle with a resolution.
type RectangleRes struct {
	Min    Point
	Max    Point
	ResLat float64
	ResLon float64
}

// LocationString returns r as a LocationString.
func (r RectangleRes) LocationString() LocationString {
	return formatFloat(r.Max.Lat) + "," + formatFloat(r.Min.Lon) +
		"_" + formatFloat(r.Min.Lat) + "," + formatFloat(r.Max.Lon) +
		":" + formatFloat(r.ResLat) + "," + formatFloat(r.ResLon)
}

// A Postal is a country code and a ZIP code.
type Postal struct {
	CountryCode string
	ZIPCode     string
}

// LocationString returns p as a LocationString.
func (p Postal) LocationString() LocationString {
	return LocationString("postal_" + p.CountryCode + p.ZIPCode)
}

// A LocationStringSlice is a slice of LocationStringers.
type LocationStringSlice []LocationStringer

// LocationString returns s as a LocationString.
func (s LocationStringSlice) LocationString() LocationString {
	ss := make([]string, len(s))
	for i, ls := range s {
		ss[i] = string(ls.LocationString())
	}
	return LocationString(strings.Join(ss, "+"))
}

func formatFloat(x float64) LocationString {
	return LocationString(strconv.FormatFloat(x, 'f', -1, 64))
}

func formatInt(i int) LocationString {
	return LocationString(strconv.Itoa(i))
}
