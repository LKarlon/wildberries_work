package engine

import (
	"fmt"
)

//интерфей для двигателя
type Engine interface {
	EngineOn()
	EngineOff()
}

// я хотел бы изменять статус двигателя (вкл, выкл),
// чтобы в зависимости от него машина могла/не могла двигаться
// но как это сделать, если реализация через интерфесы позволяет
// передавать методы, но не поля?
type engine struct {
	status bool
}

func (e *engine) EngineOn() {
	e.status = true
	fmt.Println("Двигатель включен")
}

func (e *engine) EngineOff() {
	e.status = false
	fmt.Println("Двигатель выключен")
}

func NewEngine() Engine {
	return &engine{
		status: false,
	}
}
