package factory

import "factory-method-go/pkg/vehicle"

// TruckFactory produces trucks.
type TruckFactory struct{}

// NewTruckFactory initializes a new TruckFactory.
func NewTruckFactory() *TruckFactory {
	return &TruckFactory{}
}

// CreateVehicle constructs a new Truck.
func (f *TruckFactory) CreateVehicle() vehicle.Vehicle {
	return vehicle.NewTruck("diesel", "cargo")
}
