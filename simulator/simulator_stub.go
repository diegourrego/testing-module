package simulator

import "testdoubles/positioner"

type StubSubject struct {
	// position of the subject
	Position *positioner.Position
	// speed of the subject (in m/s)
	Speed float64
}

type StubCatchSimulator interface {
	// CanCatch returns true if the hunter can catch the prey
	// - hunter: is the hunter subject
	// - prey: is the prey subject
	CanCatch(hunter, prey *Subject) (canCatch bool)
}
