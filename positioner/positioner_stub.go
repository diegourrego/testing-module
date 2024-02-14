package positioner

// Exercise #1

func NewStubPositioner() *StubPositioner {
	return &StubPositioner{}
}

type StubPositioner struct {
	FuncGetLinearDistance func(from, to *Position) (linearDistance float64)
}

func (sp *StubPositioner) GetLinearDistance(from, to *Position) (linearDistance float64) {
	linearDistance = sp.FuncGetLinearDistance(from, to)
	return
}
