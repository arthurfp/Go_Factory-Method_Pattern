package main

import (
	"factory-method-go/pkg/factory"
	"fmt"
)

func main() {
	fmt.Println("Factory Method Pattern in Go")

	// Create a car using the CarFactory
	carFactory := factory.NewCarFactory()
	car := carFactory.CreateVehicle()
	fmt.Println("Created Vehicle - Car:", car.Details())

	// Create a bike using the BikeFactory
	bikeFactory := factory.NewBikeFactory()
	bike := bikeFactory.CreateVehicle()
	fmt.Println("Created Vehicle - Bike:", bike.Details())

	// Create a truck using the TruckFactory
	truckFactory := factory.NewTruckFactory()
	truck := truckFactory.CreateVehicle()
	fmt.Println("Created Vehicle - Truck:", truck.Details())
}
