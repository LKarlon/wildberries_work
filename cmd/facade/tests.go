package main

import (
	"testing"

	"github.com/LKarlon/wildberries-work/pkg/facade/battery"
	"github.com/LKarlon/wildberries-work/pkg/facade/car"
	"github.com/LKarlon/wildberries-work/pkg/facade/engine"
	"github.com/LKarlon/wildberries-work/pkg/facade/headlights"
	"github.com/LKarlon/wildberries-work/pkg/facade/wheels"
)

func TestOk(t *testing.T){
	wheel := wheels.NewRiding()
	engin := engine.NewEngine(wheel)
	headlight := headlights.NewHeadlights()
	batter := battery.NewBattery(350, engin, headlight)
	ca := car.NewCar(headlight, engin, batter)
	okResult := 240
	res, _ := ca.Ride(5)
	if res != okResult	{
		t.Errorf("Test failed!\nExpected:\n%v\nGot %v", okResult, res)
	}
	okResult = 350
	res, _ = ca.Ride(0)
	if res != okResult	{
		t.Errorf("Test failed!\nExpected:\n%v\nGot %v", okResult, res)
	}
	okResult = 0
	res, _ = ca.Ride(-10)
	if res != okResult	{
		t.Errorf("Test failed!\nExpected:\n%v\nGot %v", okResult, res)
	}
}