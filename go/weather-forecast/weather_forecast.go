// Package weather helps with the weather forecast of various cities in Goblinocus.
package weather

// CurrentCondition represents the current weather of the location provided.
var CurrentCondition string
// CurrentLocation represents the location for which the weather has to be forecasted.
var CurrentLocation string

// Forecast function takes city and condition as string inputs and returns the forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
