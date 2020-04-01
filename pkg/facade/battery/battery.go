package battery

type engine interface {
	On(int, int) (int, error)
	Off()
}

type headlights interface {
	LampsOn(int, int) (int, error)
	LampsOff()
}

// Battery starts up the engine and headlights.
type Battery interface {
	EngineOn(int) (int, error)
	HeadlightsOn(int) (int, error)
}

type battery struct {
	engine     engine
	headlights headlights
	charge     int
}

// EngineOn starts up the engine if charge is enough
func (b *battery) EngineOn(tripLength int) (int, error) {
	balance, err := b.engine.On(b.charge, tripLength)
	if err != nil {
		return 0, err
	}
	b.charge = balance
	return b.charge, nil
}

// HeadlightsOn starts up the engine if charge is enough
func (b *battery) HeadlightsOn(tripLength int) (int, error) {
	balance, err := b.headlights.LampsOn(b.charge, tripLength)
	if err != nil {
		return 0, err
	}
	b.charge = balance
	return b.charge, nil
}

// NewBattery ...
func NewBattery(charge int, engine engine, headlights headlights) Battery {
	return &battery{
		charge:     charge,
		engine:     engine,
		headlights: headlights,
	}
}
