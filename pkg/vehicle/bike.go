package vehicle

import "fmt"

// Bike represents a type of vehicle.
type Bike struct {
	Style     string
	GearCount int
}

// NewBike creates a bike with specified style and gear count.
func NewBike(style string, gearCount int) *Bike {
	return &Bike{Style: style, GearCount: gearCount}
}

// Details provides a detailed description of the bike.
func (b *Bike) Details() string {
	return fmt.Sprintf("Bike Style: %s, Gears: %d", b.Style, b.GearCount)
}
