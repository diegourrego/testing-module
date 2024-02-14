package simulator

import (
	"github.com/stretchr/testify/require"
	"testdoubles/positioner"
	"testing"
)

func TestMockCatchSimulator_CanCatch(t *testing.T) {
	t.Run("Subject reach the other one", func(t *testing.T) {
		// Arrange
		ps := positioner.NewStubPositioner()
		ps.FuncGetLinearDistance = func(from, to *positioner.Position) (linearDistance float64) {
			linearDistance = 50
			return
		}
		cs := NewCatchSimulatorDefault(50, ps)

		// Act
		hunter := &Subject{
			Position: &positioner.Position{
				X: 2,
				Y: 1,
				Z: 1,
			},
			Speed: 20,
		}

		prey := &Subject{
			Position: &positioner.Position{
				X: 1,
				Y: 0,
				Z: 0,
			},
			Speed: 10,
		}

		obtainedCanCatch := cs.CanCatch(hunter, prey)

		// Assert
		require.Equal(t, true, obtainedCanCatch)
	})

	t.Run("Hunter is slower than prey", func(t *testing.T) {
		// Arrange
		ps := positioner.NewStubPositioner()
		ps.FuncGetLinearDistance = func(from, to *positioner.Position) (linearDistance float64) {
			linearDistance = 30
			return
		}
		cs := NewCatchSimulatorDefault(100, ps)

		// Act
		hunter := &Subject{
			Position: &positioner.Position{
				X: 2,
				Y: 1,
				Z: 1,
			},
			Speed: 20,
		}

		prey := &Subject{
			Position: &positioner.Position{
				X: 1,
				Y: 0,
				Z: 0,
			},
			Speed: 100,
		}

		obtainedCanCatch := cs.CanCatch(hunter, prey)

		// Assert
		require.Equal(t, false, obtainedCanCatch)
	})

	t.Run("Hunter can't catch the prey. Time's over", func(t *testing.T) {
		// Arrange
		ps := positioner.NewStubPositioner()
		ps.FuncGetLinearDistance = func(from, to *positioner.Position) (linearDistance float64) {
			linearDistance = 100
			return
		}
		cs := NewCatchSimulatorDefault(30, ps)

		// Act
		hunter := &Subject{
			Position: &positioner.Position{
				X: 2,
				Y: 1,
				Z: 1,
			},
			Speed: 12,
		}

		prey := &Subject{
			Position: &positioner.Position{
				X: 1,
				Y: 0,
				Z: 0,
			},
			Speed: 10,
		}

		obtainedCanCatch := cs.CanCatch(hunter, prey)

		// Assert
		require.Equal(t, false, obtainedCanCatch)
	})

}
