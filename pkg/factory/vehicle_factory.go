package factory

import "factory-method-go/pkg/vehicle"

// VehicleFactory is the interface for our factory method.
type VehicleFactory interface {
	CreateVehicle() vehicle.Vehicle
}
