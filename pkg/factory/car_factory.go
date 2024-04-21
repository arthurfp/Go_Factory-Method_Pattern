package factory

import "factory-method-go/pkg/vehicle"

// CarFactory produces cars.
type CarFactory struct{}

// NewCarFactory initializes a new CarFactory.
func NewCarFactory() *CarFactory {
	return &CarFactory{}
}

// CreateVehicle constructs a new Car.
func (f *CarFactory) CreateVehicle() vehicle.Vehicle {
	return vehicle.NewCar("sedan", "gasoline")
}
