package asteroid

import (
	"sort"

	"github.com/mfesenko/adventofcode/2019/math"
)

// Vaporizer detects the order in which the asteroids will be destroyed
type Vaporizer struct {
	asteroidMap Map
}

// NewVaporizer creates a Vaporizer
func NewVaporizer(asteroidMap Map) *Vaporizer {
	return &Vaporizer{
		asteroidMap: asteroidMap,
	}
}

// Vaporize returns locations of asteroids in the order in which they will be destroyed
func (v *Vaporizer) Vaporize(location math.Point) []math.Point {
	asteroidSlopes := v.getSlopeForAsteroids(location)
	slopes := v.getSlopes(asteroidSlopes)
	v.sortSlopes(slopes)
	return v.collectPoints(slopes, asteroidSlopes)

}

func (v *Vaporizer) getSlopeForAsteroids(location math.Point) map[slope][]math.Point {
	asteroidSlopes := map[slope][]math.Point{}
	for _, asteroid := range v.asteroidMap.Asteroids() {
		if asteroid == location {
			continue
		}
		slope := calculateSlope(asteroid, location)
		asteroidSlopes[slope] = append(asteroidSlopes[slope], asteroid)
	}

	for _, asteroids := range asteroidSlopes {
		sort.SliceStable(asteroids, func(i, j int) bool {
			a := asteroids[i]
			b := asteroids[j]
			return a.ManhattanDistance(location) < b.ManhattanDistance(location)
		})
	}
	return asteroidSlopes
}

func (v *Vaporizer) getSlopes(asteroidSlopes map[slope][]math.Point) []slope {
	slopes := make([]slope, 0)
	for slope := range asteroidSlopes {
		slopes = append(slopes, slope)
	}
	return slopes
}

func (v *Vaporizer) sortSlopes(slopes []slope) {
	sort.SliceStable(slopes, func(i, j int) bool {
		dxi := slopes[i].dx
		dxj := slopes[j].dx
		dyi := slopes[i].dy
		dyj := slopes[j].dy

		if dxi >= 0 && dxj < 0 {
			return true
		}

		if dxi < 0 && dxj >= 0 {
			return false
		}

		if dxi == 0 && dxj == 0 {
			return dyi < dyj
		}

		det := dxi*dyj - dxj*dyi
		if det <= 0 {
			return false
		}
		return true
	})
}

func (v *Vaporizer) collectPoints(slopes []slope, asteroidSlopes map[slope][]math.Point) []math.Point {
	var result []math.Point

	hasPoints := true
	for hasPoints {
		keepGoing := false
		for _, slope := range slopes {
			pointsForSlope := asteroidSlopes[slope]
			if len(pointsForSlope) > 0 {
				result = append(result, pointsForSlope[0])
				asteroidSlopes[slope] = pointsForSlope[1:]
				keepGoing = true
			}
		}
		hasPoints = keepGoing
	}

	return result
}
