package planet

import (
	"fmt"
	"strings"

	"github.com/mfesenko/adventofcode/2019/math"
)

// StateRepeater finds a step count required for the system to reach the previous state
type StateRepeater struct {
	moons []*Moon
}

// NewStateRepeater creates a new StateRepeater
func NewStateRepeater(moons []*Moon) *StateRepeater {
	return &StateRepeater{
		moons: copyMoons(moons),
	}
}

// FindStepCountForRepeat finds a step count required for the system to reach the previous state
func (r *StateRepeater) FindStepCountForRepeat() int64 {
	dimensions := r.moons[0].Position().Dimensions()
	result := int64(1)
	for _, dimension := range dimensions {
		steps := r.simulateUntilRepeatForDimension(dimension)
		result = math.LCM(steps, result)
	}
	return result
}

func (r *StateRepeater) simulateUntilRepeatForDimension(dimension string) int64 {
	simulator := NewSimulator(r.moons)
	states := map[string]struct{}{}
	i := int64(0)
	for {
		state := r.state(simulator.Moons(), dimension)
		if _, ok := states[state]; ok {
			break
		}
		states[state] = struct{}{}

		simulator.Simulate(1)
		i++
	}
	return i
}

func (r *StateRepeater) state(moons []*Moon, dimension string) string {
	state := make([]string, len(moons))
	for i, moon := range moons {
		state[i] = fmt.Sprintf("%v,%v", moon.Position().Get(dimension), moon.Velocity().Get(dimension))
	}
	return strings.Join(state, ";")
}
