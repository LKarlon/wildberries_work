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
	batter := battery.NewBattery(350, engin, headlight)
	ca := car.NewCar(headlight, engin, batter)
	_, err := ca.Ride(-10)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}
