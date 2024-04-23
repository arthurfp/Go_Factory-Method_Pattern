package factory

import (
	"factory-method-go/pkg/vehicle"
	"testing"
)

// TestBikeFactory checks if BikeFactory creates bikes correctly.
func TestBikeFactory(t *testing.T) {
	factory := NewBikeFactory()
	bike := factory.CreateVehicle()

	_, ok := bike.(*vehicle.Bike)
	if !ok {
		t.Errorf("BikeFactory did not create a Bike")
	}
}
