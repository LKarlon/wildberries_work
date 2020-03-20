package engine

import (
	"fmt"
)

type Engine struct {
}

func (e *Engine) EngineOn() {
	fmt.Println("Двигатель включен")
}

func (e *Engine) RngineOff() {
	fmt.Println("Двигатель выключен")
}
