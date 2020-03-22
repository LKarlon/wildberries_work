package riding

import (
	"fmt"
)

type Riding struct {
}

func (r *Riding) Start() {
	fmt.Println("машина движется")
}

func (r *Riding) Stop() {
	fmt.Println("мы остановились")
}
