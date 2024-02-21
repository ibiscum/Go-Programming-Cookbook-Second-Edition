package goflow

import (
	"github.com/trustmaster/goflow"
)

// NewEncodingApp wires together the components
func NewEncodingApp() *goflow.Graph {
	e := goflow.NewGraph()

	// define component types
	err := e.Add("encoder", new(Encoder))
	if err != nil {
		panic(err)
	}

	err = e.Add("printer", new(Printer))
	if err != nil {
		panic(err)
	}

	// connect the components using channels
	err = e.Connect("encoder", "Res", "printer", "Line")
	if err != nil {
		panic(err)
	}

	// map the in channel to Val, which is
	// tied to OnVal function
	e.MapInPort("In", "encoder", "Val")

	return e
}
