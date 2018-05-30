// Package space concerns things around calculating age on difference planets
// in the Solar System
package space

type Planet string

const earthYearSeconds = 31557600

// planetYears is a map of a planet to the number of earth years it experiences
// for each fo its own years.
var planetYears = map[Planet]float64{
	"Earth":   1,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age calculates the age of a period of time calculated in seconds
// for a particular planet
func Age(seconds float64, planet Planet) float64 {
	years := planetYears[planet] * earthYearSeconds
	return seconds / years
}
