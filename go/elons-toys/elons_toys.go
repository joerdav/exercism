package elon

import "fmt"

func (c *Car) Drive() {
	if c.battery < c.batteryDrain {
		return
	}
	c.distance += c.speed
	c.battery -= c.batteryDrain
}
func (c Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}
func (c Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}
func (c Car) CanFinish(trackDistance int) bool {
	for c.battery >= c.batteryDrain {
		c.Drive()
	}
	return c.distance >= trackDistance
}
