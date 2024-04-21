package factory

import "factory-method-go/pkg/vehicle"

// BikeFactory produces bikes.
type BikeFactory struct{}

// NewBikeFactory initializes a new BikeFactory.
func NewBikeFactory() *BikeFactory {
	return &BikeFactory{}
}

// CreateVehicle constructs a new Bike.
func (f *BikeFactory) CreateVehicle() vehicle.Vehicle {
	return vehicle.NewBike("mountain bike")
}
