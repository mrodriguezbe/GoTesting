package simulator_test

import (
	"gotesting/clase3/hunter"
	"gotesting/clase3/positioner"
	"gotesting/clase3/prey"
	"gotesting/clase3/simulator"
	"testing"

	"github.com/stretchr/testify/require"
)

// CatchSimulator mock
type CatchSimulatorMock struct {
	catched bool
}

func (c *CatchSimulatorMock) Catch(prey prey.Prey, hunter hunter.Hunter) (catched bool) {
	return c.catched
}

// Catch simulator test
func TestSimulation(t *testing.T) {
	t.Run("case catch simulator", func(t *testing.T) {
		p := positioner.NewPositionerDefault()
		s := simulator.NewCatchSimulatorDefault(10, p)
		ht := &simulator.Subject{
			Speed:    10,
			Position: &positioner.Position{X: 0, Y: 0, Z: 0},
		}
		pr := &simulator.Subject{
			Speed:    5,
			Position: &positioner.Position{X: 2, Y: 2, Z: 0},
		}

		result := s.CanCatch(ht, pr)

		require.True(t, result)
	})
}
