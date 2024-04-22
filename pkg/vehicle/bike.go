package vehicle

// Bike represents a type of vehicle.
type Bike struct {
	Style string
}

// NewBike creates a bike with a specified style.
func NewBike(style string) *Bike {
	return &Bike{Style: style}
}

// Details provides a description of the bike.
func (b *Bike) Details() string {
	return "Bike Style: " + b.Style
}
