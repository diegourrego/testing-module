package simulator

import "github.com/stretchr/testify/mock"

func NewMockCatchSimulator() *MockCatchSimulator {
	return &MockCatchSimulator{}
}

type MockCatchSimulator struct {
	mock.Mock
	FuncCanCatch func(hunter, prey *Subject) (canCatch bool)
}

//type MockCatchSimulator struct {
//	FuncCanCatch func(hunter, prey *Subject) (canCatch bool)
//	Calls        struct {
//		CanCatch int
//	}
//}

func (m *MockCatchSimulator) CanCatch(hunter, prey *Subject) (canCatch bool) {
	output := m.Called(hunter, prey)

	if m.FuncCanCatch != nil {
		m.FuncCanCatch(hunter, prey)
	}

	canCatch = output.Bool(0)
	return
}

//func (m *MockCatchSimulator) CanCatch(hunter, prey *Subject) (canCatch bool) {
//	m.Calls.CanCatch++
//	canCatch = m.CanCatch(hunter, prey)
//	return
//}
