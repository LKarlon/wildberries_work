package order

type Detail interface {
	Accept() string
}

type OrderSheet interface {
	Accept() string
}

type orderSheet struct {
	details []Detail
}

func (o *orderSheet) Add(d Detail) {
	o.details = append(o.details, d)
}

func (o *orderSheet) Accept() string {
	res := ""
	for _, d := range o.details {
		res += d.Accept()
	}
	return res
}

func NewOrderSheet(d ...Detail) OrderSheet {
	return &orderSheet{details: d}
}
