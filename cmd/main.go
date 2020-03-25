package main

import (
	"wildberries-work/pkg/facade/car"
	"wildberries-work/pkg/facade/engine"
	"wildberries-work/pkg/facade/headlights"
	"wildberries-work/pkg/facade/riding"
)

func main() {
	//инициализируем различные интерфесы автомобиля
	e := engine.NewEngine()
	r := riding.NewRiding()
	h := headlights.NewHeadlights()
	//компонуем из них машину
	c := car.NewRider(h, e, r)
	//вызываем функцию поездки, передав в нее время суток
	//(morning, day, evening, night) В случе некорректного ввода,
	//функция сообщит нам об этом
	c.Ride("night")
}
