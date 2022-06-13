package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

func (u TemperatureUnit) String() string {
	switch u {
	case Celsius:
		return "°C"
	case Fahrenheit:
		return "°F"
	default:
		return ""
	}
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (t Temperature) String() string {
	return fmt.Sprintf("%v %v", t.degree, t.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

func (u SpeedUnit) String() string {
	switch u {
	case MilesPerHour:
		return "mph"
	case KmPerHour:
		return "km/h"
	default:
		return ""
	}
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

func (s Speed) String() string {
	return fmt.Sprintf("%v %v", s.magnitude, s.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

func (d MeteorologyData) String() string {
	return fmt.Sprintf("%v: %v, Wind %v at %v, %v%% Humidity", d.location, d.temperature, d.windDirection, d.windSpeed, d.humidity)
}
