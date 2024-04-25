package vehicle

import "fmt"

// Car represents a type of vehicle.
type Car struct {
	CarType  string
	Engine   string
	MaxSpeed int
}

// NewCar creates a car with specified features.
func NewCar(carType, engine string, maxSpeed int) *Car {
	return &Car{CarType: carType, Engine: engine, MaxSpeed: maxSpeed}
}

// Details provides a detailed description of the car.
func (c *Car) Details() string {
	return fmt.Sprintf("Car Type: %s, Engine: %s, Max Speed: %dkm/h", c.CarType, c.Engine, c.MaxSpeed)
}
