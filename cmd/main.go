package main

import (
	"wildberries_work/pkg/facade/car"
	"wildberries_work/pkg/facade/riding"
	"wildberries_work/pkg/facade/headlights"
	"wildberries_work/pkg/facade/engine"
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
