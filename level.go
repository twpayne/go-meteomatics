package meteomatics

import "strconv"

// A LevelString is a string representation of a level.
type LevelString string

// A LevelStringer can be converted to a LevelString.
type LevelStringer interface {
	LevelString() LevelString
}

// A LevelCentimeters is level in centimeters.
type LevelCentimeters int

// LevelString returns l as a LevelString.
func (l LevelCentimeters) LevelString() LevelString {
	return LevelString(strconv.Itoa(int(l)) + "cm")
}

// A LevelMeters is level in meters.
type LevelMeters int

// LevelString returns l as a LevelString.
func (l LevelMeters) LevelString() LevelString {
	return LevelString(strconv.Itoa(int(l)) + "m")
}

// A LevelHectopascals is level in hectopascals.
type LevelHectopascals int

// LevelString returns l as a LevelString.
func (l LevelHectopascals) LevelString() LevelString {
	return LevelString(strconv.Itoa(int(l)) + "hPa")
}
