package main

import (
	"wildberries_work/pkg/facade/get_in_the_car"
)

func main() {
	car := get_in_the_car.GetInTheCar()
	engine := car.GetEngine()
	headlights := car.GetHeadLights()
	riding := car.GetRiding()
	engine.EngineOn()
	headlights.TurnOnLamps()
	riding.Start()
	riding.Stop()
	headlights.TurnOffLamps()
	engine.RngineOff()
}
