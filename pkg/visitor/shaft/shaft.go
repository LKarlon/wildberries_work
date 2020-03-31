package shaft

type cnc interface {
	SharpenShaft(sh *shaft) string
}

type Shaft interface {
	Accept() string
	VolumeCalc() int
}

type shaft struct {
	cnc
	startVolume int
	length      int
	radius      int
}

func (s *shaft) Accept() string {
	return s.cnc.SharpenShaft(s)
}

func (s *shaft) VolumeCalc() int {
	return (s.radius + 2) * 2 * s.length
}

func NewShaft(cnc cnc, startVolume int, radius int, length int) Shaft {
	return &shaft{
		cnc:         nil,
		startVolume: startVolume,
		length:      radius,
		radius:      length,
	}
}
