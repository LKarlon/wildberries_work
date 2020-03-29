package main

import (
	"fmt"
	"wildberries-work/pkg/facade/battery"
	"wildberries-work/pkg/facade/car"
	"wildberries-work/pkg/facade/engine"
	"wildberries-work/pkg/facade/headlights"
	"wildberries-work/pkg/facade/wheels"
)

func main() {

	wheels := wheels.NewRiding()
	engine := engine.NewEngine(wheels)
	headlights := headlights.NewHeadlights()
	battery := battery.NewBattery(350, engine, headlights)
	car := car.NewCar(headlights, engine, battery)
	err := car.Ride(30)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}
