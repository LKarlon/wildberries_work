package order_sheet

import(
	"github.com/LKarlon/wildberries-work/pkg/visitor/disk"
	"github.com/LKarlon/wildberries-work/pkg/visitor/shaft"
)

type Detail interface {
	Accept(c CNC) string
}

type orderSheet struct {
	details []Detail
}

func (o *orderSheet) Add(d Detail) {
	o.details = append(o.details, d)
}

func (o *orderSheet) Accept(c CNC) string {
	var res string
	for _, d := range o.details {
		res += d.Accept(c)
	}
	return res
}


