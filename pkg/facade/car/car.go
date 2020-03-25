package car

import (
	"fmt"
)

type headlights interface {
	TurnOnLamps()
	TurnOffLamps()
}

type engine interface {
	EngineOn()
	EngineOff()
}

type riding interface {
	Start()
	Stop()
}

type Rider interface {
	Ride(timesOfDay string)
}

type rider struct {
	headlights headlights
	engine     engine
	riding     riding
}

// Эта функция в зависимости от переданного аргумента (время суток) определяет - следует ли
// включать фары. Затем осуществляется поездка на автомобиле.
func (r *rider) Ride(timesOfDay string) {
	var headlightsStatus bool
	switch timesOfDay {
	case "morning":
		r.headlights.TurnOffLamps()
		headlightsStatus = false
	case "day":
		r.headlights.TurnOffLamps()
		headlightsStatus = false
	case "evening":
		r.headlights.TurnOnLamps()
		headlightsStatus = true
	case "night":
		r.headlights.TurnOnLamps()
		headlightsStatus = true
	default:
		fmt.Println("Неправильно указано время суток")
		return
	}
	if headlightsStatus {
		fmt.Println("Темно и мы включаем фары")
	} else {
		fmt.Println("Нет необходимости включать фары")
	}
	r.engine.EngineOn()
	r.riding.Start()
	r.riding.Stop()
	r.engine.EngineOff()
}

func NewRider(headlights headlights, engine engine, riding riding) Rider {
	return &rider{
		headlights: headlights,
		engine:     engine,
		riding:     riding,
	}

}
