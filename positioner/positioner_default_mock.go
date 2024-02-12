package positioner

func NewMockPositionerDefault() (positioner *MockPositionerDefault) {
	positioner = &MockPositionerDefault{}
	return
}

type MockPositionerDefault struct {
	FuncGetLinearDistance func(from, to *Position) (linearDistance float64)
	Calls                 struct {
		GetLinearDistance int
	}
}

func (p *MockPositionerDefault) GetLinearDistance(from, to *Position) (linearDistance float64) {
	p.Calls.GetLinearDistance++
	linearDistance = p.FuncGetLinearDistance(from, to)
	return
}
