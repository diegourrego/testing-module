package prey

import "testdoubles/positioner"

// Exercise #2

func NewStubPrey() *StubPrey {
	return &StubPrey{}
}

type StubPrey struct {
	// Externalize the function to be able to test it
	FuncGetSpeed    func() (speed float64)
	FuncGetPosition func() (position *positioner.Position)
}

func (p *StubPrey) GetSpeed() (speed float64) {
	speed = p.FuncGetSpeed()
	return
}

func (p *StubPrey) GetPosition() (position *positioner.Position) {
	position = p.FuncGetPosition()
	return
}
