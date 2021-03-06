package asteroid

import (
	"github.com/mfesenko/adventofcode/2019/math"
)

type slope struct {
	dx int64
	dy int64
}

func calculateSlope(a math.Point, b math.Point) slope {
	dx := a.X - b.X
	dy := a.Y - b.Y
	gcd := math.GCD(math.Abs(dx), math.Abs(dy))
	return slope{
		dx / gcd,
		dy / gcd,
	}
}
