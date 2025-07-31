package space

// Planet define that name of planets should be string
type Planet string

// Age converts a time given in seconds to Earth years depending on planet orbit period
func Age(seconds float64, planet Planet) (age float64) {
	planets := map[Planet]float64{
		"Earth":   31557600,
		"Mercury": 31557600 * 0.2408467,
		"Venus":   31557600 * 0.61519726,
		"Mars":    31557600 * 1.8808158,
		"Jupiter": 31557600 * 11.862615,
		"Saturn":  31557600 * 29.447498,
		"Uranus":  31557600 * 84.016846,
		"Neptune": 31557600 * 164.79132,
	}

	if value, ok := planets[planet]; ok {
		return seconds / value
	}

	return -1.0
}
