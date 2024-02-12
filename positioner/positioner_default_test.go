package positioner

import (
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestPositionerDefault_GetLinearDistance(t *testing.T) {
	t.Run("case 01: coordinates are negative", func(t *testing.T) {
		// Arrange
		positioner := PositionerDefault{}
		var from Position = Position{X: -4, Y: -7, Z: -2}
		var to Position = Position{X: -1, Y: -3, Z: -1}

		//dx := from.X - to.X
		//dy := from.Y - to.Y
		//dz := from.Z - to.Z

		// Act
		expectedLinearDistance := 5.099
		obtainedLinearDistance := math.Round(positioner.GetLinearDistance(&from, &to)*1000) / 1000

		// Assert
		require.Equal(t, expectedLinearDistance, obtainedLinearDistance)
	})

	t.Run("case 02: coordinates are positive", func(t *testing.T) {
		// Arrange
		positioner := PositionerDefault{}
		var from Position = Position{X: 4, Y: 7, Z: 2}
		var to Position = Position{X: 1, Y: 3, Z: 1}

		//dx := from.X - to.X
		//dy := from.Y - to.Y
		//dz := from.Z - to.Z

		// Act
		expectedLinearDistance := 5.0990195135927845
		obtainedLinearDistance := positioner.GetLinearDistance(&from, &to)

		// Assert
		require.Equal(t, expectedLinearDistance, obtainedLinearDistance)
	})

	t.Run("case 03: coordinates returns an integer", func(t *testing.T) {
		// Arrange
		positioner := PositionerDefault{}
		var from Position = Position{X: 0, Y: 0, Z: 0}
		var to Position = Position{X: 3, Y: 4, Z: 0}

		// Act
		expectedLinearDistance := 5.0
		obtainedLinearDistance := positioner.GetLinearDistance(&from, &to)

		// Assert
		require.Equal(t, expectedLinearDistance, obtainedLinearDistance)
	})
}
