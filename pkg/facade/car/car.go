package car

import (
	"fmt"
)

type headlights interface {
	LampsOn(int, int)(int, error)
	LampsOff()
}

type engine interface {
	On(int, int) (int, error)
	Off()
	WheelsStart()
	WheelsStop()
}

type battery interface {
	EngineOn(int) (int, error)
	HeadlightsOn(int) (int, error)
}

type Car interface {
	Ride(tripLength int) error
}

type car struct {
	headlights headlights
	engine     engine
	battery    battery
}

func (c *car) Ride(tripLength int) error {
	_, err := c.battery.HeadlightsOn(tripLength)
	if err != nil {
		return err
	}
	charge, err := c.battery.EngineOn(tripLength)
	if err != nil {
		return err
	}
	c.engine.WheelsStart()
	c.engine.WheelsStop()
	c.engine.Off()
	c.headlights.LampsOff()
	fmt.Println("Осталось заряда: ", charge)
	return err
}

func NewCar(headlights headlights, engine engine, battery battery) Car {
	return &car{
		headlights: headlights,
		engine:     engine,
		battery:    battery,
	}

}
