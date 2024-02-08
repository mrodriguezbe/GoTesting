package positioner_test

import (
	"gotesting/clase2/hunter"
	"gotesting/clase2/positioner"
	"gotesting/clase2/prey"
	"testing"

	"github.com/stretchr/testify/require"
)

// Positioner stub
type PositionerStub struct{}

func (p *PositionerStub) GetLinearDistance(from, to *positioner.Position) (linearDistance float64) {
	// need implementation
	return 5.0
}

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

// CatchSimulator mock
type CatchSimulatorMock struct {
	catched bool
}

func (c *CatchSimulatorMock) Catch(prey prey.Prey, hunter hunter.Hunter) (catched bool) {
	return c.catched
}

// TestPositioner get linear distance
func TestPositioner(t *testing.T) {
	t.Run("case all coordinates are negative", func(t *testing.T) {
		p := positioner.NewPositionerDefault()
		start := &positioner.Position{X: -1, Y: -1, Z: -1}
		end := &positioner.Position{X: -2, Y: -2, Z: -2}

		result := p.GetLinearDistance(start, end)

		require.Equal(t, 1.7320508075688772, result)
	})

	t.Run("case all coordinates are positive", func(t *testing.T) {
		p := positioner.NewPositionerDefault()
		start := &positioner.Position{X: 1, Y: 1, Z: 1}
		end := &positioner.Position{X: 2, Y: 2, Z: 2}

		result := p.GetLinearDistance(start, end)

		require.Equal(t, 1.7320508075688772, result)
	})

	t.Run("case the coordinates return a positive lineal distance without decimal", func(t *testing.T) {
		p := positioner.NewPositionerDefault()
		start := &positioner.Position{X: 1, Y: 0, Z: 0}
		end := &positioner.Position{X: 2, Y: 0, Z: 0}

		result := p.GetLinearDistance(start, end)

		require.Equal(t, 1.0, result)
	})
}
