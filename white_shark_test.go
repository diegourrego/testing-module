package hunt_test

import (
	"github.com/stretchr/testify/require"
	hunt "testdoubles"
	"testing"
)

// Tests for the WhiteShark implementation - Hunt method
func TestWhiteSharkHunt(t *testing.T) {
	t.Run("case 1: white shark hunts successfully", func(t *testing.T) {
		// Arrange
		shark := hunt.NewWhiteShark(true, false, 237.6)
		tuna := hunt.NewTuna("my test tuna", 200.03)
		// Act
		err := shark.Hunt(tuna)
		// Assert
		require.NoError(t, err)
		require.Equal(t, false, shark.Hungry)
		require.Equal(t, true, shark.Tired)
	})

	t.Run("case 2: white shark is not hungry", func(t *testing.T) {
		// Arrange
		shark := hunt.NewWhiteShark(false, false, 237.6)
		tuna := hunt.NewTuna("my test tuna", 200.03)
		// Act
		err := shark.Hunt(tuna)
		// Assert
		require.Error(t, err, "can not hunt, shark is not hungry")
		require.Equal(t, false, shark.Hungry)
		require.Equal(t, false, shark.Tired)
	})

	t.Run("case 3: white shark is tired", func(t *testing.T) {
		// Arrange
		shark := hunt.NewWhiteShark(true, true, 237.6)
		tuna := hunt.NewTuna("my test tuna", 200.03)
		// Act
		err := shark.Hunt(tuna)
		// Assert
		require.Error(t, err, "can not hunt, shark is tired")
		require.Equal(t, true, shark.Hungry)
		require.Equal(t, true, shark.Tired)
	})

	t.Run("case 4: white shark is slower than the tuna", func(t *testing.T) {
		// Arrange
		shark := hunt.NewWhiteShark(true, false, 123.9)
		tuna := hunt.NewTuna("my test tuna", 200.03)
		// Act
		err := shark.Hunt(tuna)
		// Assert
		require.Error(t, err, "can not hunt, shark is slower than the prey")
		require.Equal(t, true, shark.Hungry)
		require.Equal(t, false, shark.Tired)
	})

	t.Run("case 5: tuna is nil", func(t *testing.T) {
		// Arrange
		shark := hunt.NewWhiteShark(true, false, 123.9)
		// Act
		err := shark.Hunt(nil)
		// Assert
		require.Error(t, err, "can not hunt, tuna was no provided")
		require.Equal(t, true, shark.Hungry)
		require.Equal(t, false, shark.Tired)
	})
}
