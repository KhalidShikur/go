package jedlik
import "fmt"

// TODO: define the 'Drive()' method
func (c *Car) Drive() {
    if c.batteryDrain <= c.battery {
        c.distance += c.speed
        c.battery -= c.batteryDrain
    }
    return
}

// TODO: define the 'DisplayDistance() string' method
func (c Car) DisplayDistance() string {
    return "Driven " + fmt.Sprintf("%d", c.distance) + " meters"
}

// TODO: define the 'DisplayBattery() string' method
func (c Car) DisplayBattery() string {
    return "Battery at " + fmt.Sprintf("%d", c.battery) + "%"
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (c Car) CanFinish(trackDistance int) bool {
    totalDistance := (int(c.battery/c.batteryDrain) * c.speed) + c.distance
    return totalDistance >= trackDistance
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
