package get_in_the_car

type headlights interface {
	TurnOnLamps()
	TurnOffLamps()
}

type engine interface {
	EngineOn()
	EngineOff()
}

type riding interface {
	Start()
	Stop()
}

type Rider interface {
	Ride(timesOfDay string) string
}

type rider struct {
	headlights headlights
	engine     engine
	riding     riding
}

func (r *rider) Ride(timesOfDay string) (status string){
	switch timesOfDay {
	case "morning":
		r.headlights.TurnOffLamps()
	case "day":
		r.headlights.TurnOffLamps()
	case "evening":
		r.headlights.TurnOnLamps()
	case "night":
		r.headlights.TurnOnLamps()

	}
}
type road struct {
	headlights: &headlights.headlights{
},
	engine:     &engine.Engine{
},
	riding:     &riding.Riding{
},
}

func (r *sportCar) GetHeadLights() headlights.headlights {
	return *r.headlights
}

func (r *sportCar) GetEngine() engine.Engine {
	return *r.engine
}

func (r *sportCar) GetRiding() riding.Riding {
	return *r.riding
}

func GetInTheCar() *sportCar {
	return &sportCar{
		headlights: &headlights.headlights{},
		engine:     &engine.Engine{},
		riding:     &riding.Riding{},
	}
}
