// Package weather is a package for weather forcasting of the goblin.
package weather

var (
    // CurrentCondition is the current condition of a city.
	CurrentCondition string
    // CurrentLocation is the city whom it's weather is going to be forcasted.
	CurrentLocation  string
)

// Forecast function returns a string of the current condition of a city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
