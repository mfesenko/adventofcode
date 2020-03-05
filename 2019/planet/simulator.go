package planet

// Simulator simulates the motion of the moons
type Simulator struct {
	moons []*Moon
}

// NewSimulator creates a Simulator
func NewSimulator(moons []*Moon) *Simulator {
	return &Simulator{
		moons: copyMoons(moons),
	}
}

// Simulate runs a simulation for given amount of time steps
func (s *Simulator) Simulate(steps int) {
	for i := 0; i < steps; i++ {
		s.applyGravity()
		s.applyVelocity()
	}
}

func (s *Simulator) applyGravity() {
	for i, moon := range s.moons {
		for j := i + 1; j < len(s.moons); j++ {
			moon.ApplyGravity(s.moons[j])
		}
	}
}

func (s *Simulator) applyVelocity() {
	for _, moon := range s.moons {
		moon.ApplyVelocity()
	}
}

// TotalEnergy returns the total energy of the system
func (s *Simulator) TotalEnergy() int64 {
	energy := int64(0)
	for _, moon := range s.moons {
		energy += moon.TotalEnergy()
	}
	return energy
}

// Moons returns the current state of the moons
func (s *Simulator) Moons() []*Moon {
	return s.moons
}
