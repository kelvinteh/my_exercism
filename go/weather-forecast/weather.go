// Package weather provides functionality for weather forecasting
package weather

// CurrentCondition stores the weather condition currently
var CurrentCondition string

// CurrentLocation stores the location for the forecase
var CurrentLocation string

// Forecast returns the forecast for the location in a string when it is given the city and the current condition
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
