package factory

import (
	"factory-method-go/pkg/vehicle"
	"testing"
)

// TestTruckFactory checks if TruckFactory creates trucks correctly.
func TestTruckFactory(t *testing.T) {
	factory := NewTruckFactory()
	truck := factory.CreateVehicle()

	_, ok := truck.(*vehicle.Truck)
	if !ok {
		t.Errorf("TruckFactory did not create a Truck")
	}
}
