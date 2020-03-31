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
	WheelsStart() error
	WheelsStop()
}

type engine struct {
	status bool
	wheels wheels
}

// The engine start function takes charge amount and travel time.
// Returns the remaining amount of charge, or reports an error.
func (e *engine) On(chargeCalc int, tripLength int) (int, error) {
	spending := tripLength * 20
	if chargeCalc-spending <= 0 {
		err := fmt.Errorf("заряда не достаточно для поездки")
		return 0, err
	}
	e.status = true
	fmt.Println("Двигатель включен")
	return chargeCalc - spending, nil
}

func (e *engine) Off() {
	e.status = false
	fmt.Println("Двигатель выключен")
}

// Wheels start only when the engine is running
func (e *engine) WheelsStart() error {
	if e.status == false {
		return fmt.Errorf("необходимо включить двигатель")
	}
	e.wheels.Start()
	return nil
}

func (e *engine) WheelsStop() {
	e.wheels.Stop()
}

func NewEngine(wheels wheels) Engine {
	return &engine{
		wheels: wheels,
	}
}
