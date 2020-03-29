package headlights

import "fmt"

// интерфейс для фар
type Headlights interface {
	LampsOn(int, int)(int, error)
	LampsOff()
}

type headlights struct {
}

func (h *headlights) LampsOn(charge int, tripLength int) (int, error) {
	spending := tripLength * 2
	if charge - spending <= 0{
		err := fmt.Errorf("заряда не достаточно для поездки")
		return 0, err
	}
	fmt.Println("Фары включены")
	return charge - spending, nil
}

func (h *headlights) LampsOff() {
	fmt.Println("Фары выключены")
}

func NewHeadlights() Headlights {
	return &headlights{}
}
