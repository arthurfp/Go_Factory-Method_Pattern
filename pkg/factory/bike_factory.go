package factory

import "factory-method-go/pkg/vehicle"

// BikeFactory is a concrete factory that produces bike objects.
type BikeFactory struct{}

// NewBikeFactory creates a new instance of BikeFactory.
func NewBikeFactory() *BikeFactory {
	return &BikeFactory{}
}

// CreateVehicle returns a new Bike instance.
func (f *BikeFactory) CreateVehicle() vehicle.Vehicle {
	return vehicle.NewBike("mountain bike", 21)
}
