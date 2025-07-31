package space

// Planet define that name of planets should be string
type Planet string

const earthSeconds float64 = 31557600

var planets = map[Planet]float64{
	"Earth":   earthSeconds,
	"Mercury": 0.2408467 * earthSeconds,
	"Venus":   0.61519726 * earthSeconds,
	"Mars":    1.8808158 * earthSeconds,
	"Jupiter": 11.862615 * earthSeconds,
	"Saturn":  29.447498 * earthSeconds,
	"Uranus":  84.016846 * earthSeconds,
	"Neptune": 164.79132 * earthSeconds,
}

// Age converts a time given in seconds to Earth years depending on planet orbit period
func Age(seconds float64, planet Planet) (age float64) {

	if value, ok := planets[planet]; ok {
		return seconds / value
	}

	return -1.0
}
