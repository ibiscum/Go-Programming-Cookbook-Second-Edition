package kafkaflow

import "github.com/trustmaster/goflow"

// NewUpperApp wires together the compoents
func NewUpperApp() *goflow.Graph {
	u := goflow.NewGraph()

	err := u.Add("upper", new(Upper))
	if err != nil {
		panic(err)
	}

	err = u.Add("printer", new(Printer))
	if err != nil {
		panic(err)
	}

	err = u.Connect("upper", "Res", "printer", "Line")
	if err != nil {
		panic(err)
	}

	u.MapInPort("In", "upper", "Val")

	return u
}
