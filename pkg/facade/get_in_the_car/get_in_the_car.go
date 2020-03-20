package get_in_the_car

import (
	"wildberries_work/pkg/facade/engine"
	"wildberries_work/pkg/facade/headlights"
	"wildberries_work/pkg/facade/riding"
)

type SportCar struct {
	headlights *headlights.Headlights
	engine     *engine.Engine
	riding     *riding.Riding
}

func (r *SportCar) GetHeadLights() headlights.Headlights {
	return *r.headlights
}

func (r *SportCar) GetEngine() engine.Engine {
	return *r.engine
}

func (r *SportCar) GetRiding() riding.Riding {
	return *r.riding
}

func GetInTheCar() *SportCar {
	return &SportCar{
		headlights: &headlights.Headlights{},
		engine:     &engine.Engine{},
		riding:     &riding.Riding{},
	}
}
