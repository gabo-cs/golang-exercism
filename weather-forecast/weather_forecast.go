// Package weather implements a simple weather forecast functionality.
package weather

// CurrentCondition stores the weather condition.
var CurrentCondition string

// CurrentLocation stores the city to give the forecast about.
var CurrentLocation string

// Forecast indicates the current weather condition of a specific city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
