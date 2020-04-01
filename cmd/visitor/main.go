package main

import(
	"fmt"

	"github.com/LKarlon/wildberries-work/pkg/visitor/shaft"
	"github.com/LKarlon/wildberries-work/pkg/visitor/disk"
	"github.com/LKarlon/wildberries-work/pkg/visitor/cnc"
	"github.com/LKarlon/wildberries-work/pkg/visitor/order"
)

func main(){
	nc := cnc.NewCNC()
	shaf := shaft.NewShaft(nc, 500, 20, 50)
	dis := disk.NewDisk(nc,500, 20, 5)
	orde := order.NewOrderSheet(shaf, dis)
	res := orde.Accept()
	fmt.Println(res)
}
