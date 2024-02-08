package hunt_test

import (
	hunt "gotesting/clase1"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for the WhiteShark implementation - Hunt method
func TestWhiteSharkHunt(t *testing.T) {
	t.Run("case 1: white shark hunts successfully", func(t *testing.T) {
		shark := hunt.NewWhiteShark(true, false, 10.0)
		tuna := hunt.NewTuna("Tunita", 9)

		err := shark.Hunt(tuna)

		require.NoError(t, err)
		require.False(t, shark.Hungry)
		require.True(t, shark.Tired)
	})

	t.Run("case 2: white shark is not hungry", func(t *testing.T) {
		shark := hunt.NewWhiteShark(true, false, 10.0)
		tuna := hunt.NewTuna("Tunita", 9)

		shark.Hunt(tuna)

		err := shark.Hunt(tuna)
		require.Error(t, err)
		require.ErrorIs(t, err, hunt.ErrSharkIsNotHungry)
	})

	t.Run("case 3: white shark is tired", func(t *testing.T) {
		shark := hunt.NewWhiteShark(true, true, 10.0)
		tuna := hunt.NewTuna("Tunita", 19)

		shark.Hunt(tuna)

		err := shark.Hunt(tuna)
		require.Error(t, err)
		require.ErrorIs(t, err, hunt.ErrSharkIsTired)
	})

	t.Run("case 4: white shark is slower than the tuna", func(t *testing.T) {
		shark := hunt.NewWhiteShark(true, false, 10.0)
		tuna := hunt.NewTuna("Tunita", 19)

		shark.Hunt(tuna)

		err := shark.Hunt(tuna)
		require.Error(t, err)
		require.ErrorIs(t, err, hunt.ErrSharkIsSlower)
	})

	t.Run("case 5: tuna is nil", func(t *testing.T) {
		shark := hunt.NewWhiteShark(true, false, 10.0)

		err := shark.Hunt(nil)

		require.Error(t, err)
		require.ErrorIs(t, err, hunt.ErrSharkTunaCannotBeNil)
	})
}
