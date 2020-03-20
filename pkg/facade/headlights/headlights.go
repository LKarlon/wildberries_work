package headlights

import (
	"fmt"
)

type Headlights struct {
}

func (р *Headlights) TurnOnLamps() {
		fmt.Println("фары включены")
}

func (h *Headlights) TurnOffLamps() {
	fmt.Println("фары выключены")
}


