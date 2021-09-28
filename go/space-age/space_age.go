package space

// Planet type from a string
type Planet string

// Age is a function
func Age(sec float64, p Planet) float64 {

	orbitalPeriods := make(map[Planet]float64)

	orbitalPeriods["Mercury"] = 0.2408467
	orbitalPeriods["Venus"] = 0.61519726
	orbitalPeriods["Earth"] = 1.0
	orbitalPeriods["Mars"] = 1.8808158
	orbitalPeriods["Jupiter"] = 11.862615
	orbitalPeriods["Saturn"] = 29.447498
	orbitalPeriods["Uranus"] = 84.016846
	orbitalPeriods["Neptune"] = 164.79132

	earthYearSecs := float64(31557600)

	return sec / (orbitalPeriods[p] * earthYearSecs)
}
