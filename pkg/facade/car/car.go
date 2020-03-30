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
	WheelsStart() error
	WheelsStop()
}

type battery interface {
	EngineOn(int) (int, error)
	HeadlightsOn(int) (int, error)
}

type Car interface {
	Ride(tripLength int) (int, error)
}

type car struct {
	headlights headlights
	engine     engine
	battery    battery
}

func (c *car) Ride(tripLength int) (int, error) {
	if tripLength < 0{
		return 0, fmt.Errorf("продолжительность поездки не может быть отрицательной")
	}
	_, err := c.battery.HeadlightsOn(tripLength)
	if err != nil {
		return 0, err
	}
	charge, err := c.battery.EngineOn(tripLength)
	if err != nil {
		return 0, err
	}
	c.engine.WheelsStart()
	c.engine.WheelsStop()
	c.engine.Off()
	c.headlights.LampsOff()
	fmt.Println("Осталось заряда: ", charge)
	return 0, err
}

func NewCar(headlights headlights, engine engine, battery battery) Car {
	return &car{
		headlights: headlights,
		engine:     engine,
		battery:    battery,
	}

}
