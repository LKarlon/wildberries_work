package riding

import "fmt"

//интерфейс для движения
type Riding interface {
	Start()
	Stop()
}

type riding struct {
	status bool
}

func (r *riding) Start() {
	r.status = true
	fmt.Println("Машина движется")
}

func (r *riding) Stop() {
	r.status = false
	fmt.Println("Машина остановлена")
}

func NewRiding() Riding {
	return &riding{
		status: false,
	}
}
