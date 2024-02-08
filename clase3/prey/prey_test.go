package prey_test

import (
	"gotesting/clase3/positioner"
	"gotesting/clase3/prey"
	"testing"

	"github.com/stretchr/testify/require"
)

// Pray stub
type PrayStub struct {
	speed    float64
	position *positioner.Position
}

func (p *PrayStub) GetSpeed() (speed float64) {
	return p.speed
}

func (p *PrayStub) GetPosition() (position *positioner.Position) {
	return p.position
}

func TestTuna(t *testing.T) {
	t.Run("get speed test", func(t *testing.T) {
		tuna := prey.NewTuna(500, &positioner.Position{X: 1, Y: 1, Z: 1})
		require.Equal(t, 500.0, tuna.GetSpeed())
	})

	t.Run("get position test", func(t *testing.T) {
		tuna := prey.NewTuna(500, &positioner.Position{X: 1, Y: 1, Z: 1})
		require.Equal(t, &positioner.Position{X: 1, Y: 1, Z: 1}, tuna.GetPosition())
	})
}
