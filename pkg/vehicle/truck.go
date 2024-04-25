package vehicle

import "fmt"

// Truck represents a type of vehicle.
type Truck struct {
	EngineType   string
	Usage        string
	LoadCapacity int // in kilograms
}

// NewTruck creates a truck with specified engine type, usage, and load capacity.
func NewTruck(engineType, usage string, loadCapacity int) *Truck {
	return &Truck{EngineType: engineType, Usage: usage, LoadCapacity: loadCapacity}
}

// Details provides a detailed description of the truck.
func (t *Truck) Details() string {
	return fmt.Sprintf("Truck Engine Type: %s, Usage: %s, Load Capacity: %dkg", t.EngineType, t.Usage, t.LoadCapacity)
}
