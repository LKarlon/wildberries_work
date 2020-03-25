package headlights

// интерфейс для фар
type Headlights interface {
	TurnOnLamps()
	TurnOffLamps()
}

type headlights struct {
	status bool
}

func (h *headlights) TurnOnLamps() {
	h.status = true
}

func (h *headlights) TurnOffLamps() {
	h.status = false
}

func NewHeadlights() Headlights {
	return &headlights{
		status: false,
	}
}
