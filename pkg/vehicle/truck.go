package vehicle

// Truck represents a type of vehicle.
type Truck struct {
	EngineType string
	Usage      string
}

// NewTruck creates a truck with specified engine type and usage.
func NewTruck(engineType, usage string) *Truck {
	return &Truck{EngineType: engineType, Usage: usage}
}

// Details provides a description of the truck.
func (t *Truck) Details() string {
	return "Truck Engine Type: " + t.EngineType + ", Usage: " + t.Usage
}
