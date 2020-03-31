package cnc

import(
	"github.com/LKarlon/wildberries-work/pkg/visitor/disk"
	"github.com/LKarlon/wildberries-work/pkg/visitor/shaft"
)

type CNC interface {
	SharpenShaft(shaft.Shaft) string
	SharpenDisk(disk.Disk) string
}

type cnc struct {
}

func (c *cnc) SharpenShaft(shaft shaft.Shaft) (res string) {
	res = "shaft ready"
	return
}

func (c *cnc) SharpenDisk(disk disk.Disk) (res string) {
	res = "disk ready"
}

func NewCNC() CNC {
	return &cnc{}
}






