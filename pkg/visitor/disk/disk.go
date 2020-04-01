package disk

type cnc interface {
	SharpenDisk(Disk) string
}

type Disk interface {
	Accept() string
	VolumeCalc() int
}

type disk struct {
	cnc
	startVolume int
	radius      int
	depth       int
}

func (d *disk) Accept() string {
	return d.cnc.SharpenDisk(d)
}

func (d *disk) VolumeCalc() int {
	return (d.radius + 2) * 2 * d.depth
}

func NewDisk(cnc cnc, startVolume int, radius int, depth int) Disk {
	return &disk{
		cnc:         cnc,
		startVolume: startVolume,
		radius:      radius,
		depth:       depth,
	}
}
