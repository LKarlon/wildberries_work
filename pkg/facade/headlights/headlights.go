package headlights

import "fmt"

type Headlights interface {
	LampsOn(int, int) (int, error)
	LampsOff()
}

type headlights struct {
	status bool
}

// Function LampsOn takes charge amount and travel time.
// Returns the remaining amount of charge, or reports an error.
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

func (h *headlights) LampsOff() {
	if h.status{
		return
	}
	fmt.Println("Фары выключены")
	h.status = false
}

func NewHeadlights() Headlights {
	return &headlights{}
}
