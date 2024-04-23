package factory

import (
	"factory-method-go/pkg/vehicle"
	"testing"
)

// TestCarFactory checks if CarFactory creates cars correctly.
func TestCarFactory(t *testing.T) {
	factory := NewCarFactory()
	car := factory.CreateVehicle()

	_, ok := car.(*vehicle.Car)
	if !ok {
		t.Errorf("CarFactory did not create a Car")
	}
}
