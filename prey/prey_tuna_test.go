package prey

import (
	"github.com/stretchr/testify/require"
	"testdoubles/positioner"
	"testing"
)

func TestTuna_GetSpeed(t *testing.T) {
	t.Run("Tuna has default speed value", func(t *testing.T) {
		// Arrange
		tuna := NewTuna(0, nil)
		expectedSpeed := 0.0
		// Act
		obtainedSpeed := tuna.GetSpeed()
		// Assert
		require.Equal(t, expectedSpeed, obtainedSpeed)
	})

	t.Run("Tuna has a different speed value", func(t *testing.T) {
		// Arrange
		tuna := NewTuna(20, nil)
		expectedSpeed := 20.0
		// Act
		obtainedSpeed := tuna.GetSpeed()
		// Assert
		require.Equal(t, expectedSpeed, obtainedSpeed)
	})
}

func TestTuna_GetPosition(t *testing.T) {
	t.Run("Position must be nil", func(t *testing.T) {
		// Arrange
		tuna := NewTuna(20, nil)
		// Act
		obtainedPosition := tuna.GetPosition()
		// Assert
		require.Nil(t, obtainedPosition)
	})

	t.Run("Position is different to nil", func(t *testing.T) {
		position := &positioner.Position{
			X: 2,
			Y: 3,
			Z: 5,
		}
		tuna := NewTuna(20, position)
		// Act
		obtainedPosition := tuna.GetPosition()
		// Assert
		require.Equal(t, position, obtainedPosition)
	})
}
