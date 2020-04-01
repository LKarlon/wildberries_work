package headlights

import "fmt"

// Headlights...
type Headlights interface {
	LampsOn(int, int) (int, error)
	LampsOff()
}

type headlights struct {
	status bool
}

// LampsOn starts up the lamps if charge is enough
func (h *headlights) LampsOn(charge int, tripLength int) (int, error) {
	spending := tripLength * 2
	if charge-spending <= 0 {
		err := fmt.Errorf("заряда не достаточно для поездки")
		return 0, err
	}
	h.status = true
	fmt.Println("Фары включены")
	return charge - spending, nil
}

// LampsOff turns off the headlights
func (h *headlights) LampsOff() {
	if h.status == false {
		return
	}
	h.status = false
	fmt.Println("Фары выключены")
}

// NewHeadlights...
func NewHeadlights() Headlights {
	return &headlights{}
}
