package planet

type testData struct {
	input       []string
	steps       int
	totalEnergy int64
	repeatSteps int64
}

func tests() []testData {
	return []testData{
		{
			input: []string{
				"<x=-1, y=0, z=2>",
				"<x=2, y=-10, z=-7>",
				"<x=4, y=-8, z=8>",
				"<x=3, y=5, z=-1>",
			},
			steps:       10,
			totalEnergy: 179,
			repeatSteps: 2772,
		},
		{
			input: []string{
				"<x=-8, y=-10, z=0>",
				"<x=5, y=5, z=10>",
				"<x=2, y=-7, z=3>",
				"<x=9, y=-8, z=-3>",
			},
			steps:       100,
			totalEnergy: 1940,
			repeatSteps: 4686774924,
		},
	}
}
