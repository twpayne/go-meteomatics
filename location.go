package meteomatics

import (
	"strconv"
	"strings"
)

type LocationString string

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

func (s LocationString) LocationString() LocationString {
	return s
}

type Point struct {
	Lat float64
	Lon float64
}

func (p Point) LocationString() LocationString {
	return formatFloat(p.Lat) + "," + formatFloat(p.Lon)
}

type PointList []Point

func (l PointList) LocationString() LocationString {
	ss := make([]string, len(l))
	for i, p := range l {
		ss[i] = string(p.LocationString())
	}
	return LocationString(strings.Join(ss, "+"))
}

type Line struct {
	Start Point
	End   Point
	N     int
}

func (l Line) LocationString() LocationString {
	return l.Start.LocationString() +
		"_" + l.End.LocationString() +
		":" + formatInt(l.N)
}

type PolylineSegment struct {
	End Point
	N   int
}

type Polyline struct {
	Start    Point
	Segments []PolylineSegment
}

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

type RectangleN struct {
	Min  Point
	Max  Point
	NLon int
	NLat int
}

func (r RectangleN) LocationString() LocationString {
	return formatFloat(r.Max.Lat) + "," + formatFloat(r.Min.Lon) +
		"_" + formatFloat(r.Min.Lat) + "," + formatFloat(r.Max.Lon) +
		":" + formatInt(r.NLon) + "x" + formatInt(r.NLat)
}

type RectangleRes struct {
	Min    Point
	Max    Point
	ResLat float64
	ResLon float64
}

func (r RectangleRes) LocationString() LocationString {
	return formatFloat(r.Max.Lat) + "," + formatFloat(r.Min.Lon) +
		"_" + formatFloat(r.Min.Lat) + "," + formatFloat(r.Max.Lon) +
		":" + formatFloat(r.ResLat) + "," + formatFloat(r.ResLon)
}

type Postal struct {
	CountryCode string
	ZIPCode     string
}

func (p Postal) LocationString() LocationString {
	return LocationString("postal_" + p.CountryCode + p.ZIPCode)
}

type LocationStringSlice []LocationStringer

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
