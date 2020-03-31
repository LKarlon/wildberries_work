package disk

type cnc interface {
	SharpenDisk(di *disk) string
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

func NewDisk(startVolume int, radius int, depth int) Disk{
	return &disk{
		startVolume: startVolume,
		radius: radius,
		depth: depth,
	}
}