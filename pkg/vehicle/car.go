package vehicle

// Car represents a type of vehicle.
type Car struct {
	CarType string
	Engine  string
}

// NewCar creates a car with specified features.
func NewCar(carType, engine string) *Car {
	return &Car{CarType: carType, Engine: engine}
}

// Details provides a description of the car.
func (c *Car) Details() string {
	return "Car Type: " + c.CarType + ", Engine: " + c.Engine
}
