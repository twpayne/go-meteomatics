package meteomatics

// PrecipType is a type of precipitation.
type PrecipType int

// Precipitation types.
const (
	PrecipNone             PrecipType = 0
	PrecipRain             PrecipType = 1
	PrecipRainAndSnowMixed PrecipType = 2
	PrecipSnow             PrecipType = 3
	PrecipSleet            PrecipType = 4
	PrecipFreezingRain     PrecipType = 5
)
