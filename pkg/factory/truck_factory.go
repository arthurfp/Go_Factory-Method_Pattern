package factory

import "factory-method-go/pkg/vehicle"

// TruckFactory is a concrete factory that produces truck objects.
type TruckFactory struct{}

// NewTruckFactory creates a new instance of TruckFactory.
func NewTruckFactory() *TruckFactory {
	return &TruckFactory{}
}

// CreateVehicle returns a new Truck instance.
func (f *TruckFactory) CreateVehicle() vehicle.Vehicle {
	return vehicle.NewTruck("diesel", "cargo", 5000)
}
