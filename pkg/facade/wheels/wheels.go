package wheels

import "fmt"

type Wheels interface {
	Start()
	Stop()
}

type wheels struct {
}

func (w *wheels) Start() {
	fmt.Println("Машина движется")
}

func (w *wheels) Stop() {
	fmt.Println("Машина остановлена")
}

func NewRiding() Wheels {
	return &wheels{
	}
}
