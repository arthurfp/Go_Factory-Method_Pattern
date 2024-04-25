package factory

import "factory-method-go/pkg/vehicle"

// CarFactory is a concrete factory that produces car objects.
type CarFactory struct{}

// NewCarFactory creates a new instance of CarFactory.
func NewCarFactory() *CarFactory {
	return &CarFactory{}
}

// CreateVehicle returns a new Car instance.
func (f *CarFactory) CreateVehicle() vehicle.Vehicle {
	return vehicle.NewCar("sedan", "gasoline", 220)
}
