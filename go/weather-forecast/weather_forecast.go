// Package weather provides a function for displaying the weather forecast.
package weather

// CurrentCondition stores the current weather condition.
var CurrentCondition string

// CurrentLocation stores the current city.
var CurrentLocation string

// Forecast returns a formatted string that states the current weather condition in a given city.
// It also updates CurrentCondition and CurrentLocation.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
