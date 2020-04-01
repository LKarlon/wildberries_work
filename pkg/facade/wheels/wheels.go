package wheels

import "fmt"

// The essence of the wheel simply displays the log on the screen.
// I added it to learn how to embed interfaces
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
