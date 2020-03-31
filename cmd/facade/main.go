package main

import (
	"fmt"

	"github.com/LKarlon/wildberries-work/pkg/facade/battery"
	"github.com/LKarlon/wildberries-work/pkg/facade/car"
	"github.com/LKarlon/wildberries-work/pkg/facade/engine"
	"github.com/LKarlon/wildberries-work/pkg/facade/headlights"
	"github.com/LKarlon/wildberries-work/pkg/facade/wheels"
)

func main() {
	wheel := wheels.NewRiding()
	engin := engine.NewEngine(wheel)
	headlight := headlights.NewHeadlights()
	// Initialize battery with charge level
	batter := battery.NewBattery(350, engin, headlight)
	ca := car.NewCar(headlight, engin, batter)
	// You must specify the travel time in the argument
	_, err := ca.Ride(5)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}
