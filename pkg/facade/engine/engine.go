package engine

import (
	"fmt"
)

type wheels interface {
	Start()
	Stop()
}

type Engine interface {
	On(int, int) (int, error)
	Off()
	WheelsStart()
	WheelsStop()
}

type engine struct {
	status bool
	wheels wheels
}

func (e *engine) On(chargeCalc int, tripLength int) (int, error) {
	spending := tripLength * 20
	if chargeCalc-spending <= 0 {
		err := fmt.Errorf("заряда не достаточно для поездки")
		return 0, err
	}
	fmt.Println("Двигатель включен")
	return chargeCalc - spending, nil
}

func (e *engine) Off() {
	fmt.Println("Двигатель выключен")
}

func (e*engine) WheelsStart(){
	e.wheels.Start()
}

func(e *engine) WheelsStop(){
	e.wheels.Stop()
}

func NewEngine(wheels wheels) Engine {
	return &engine{
		wheels: wheels,
	}
}
