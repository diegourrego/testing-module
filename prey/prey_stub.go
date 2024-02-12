package prey

import "testdoubles/positioner"

type StubPrey interface {
	// GetSpeed returns the speed of the prey
	GetSpeed() (speed float64)
	// GetPosition returns the position of the prey
	GetPosition() (position *positioner.Position)
}
