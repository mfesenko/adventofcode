package asteroid

import (
	"github.com/mfesenko/adventofcode/2019/math"
)

// Detector analyses the asteroid map in order to find the best location to build an asteroid detection station
type Detector struct {
	asteroidMap Map
}

// NewDetector creates a new Detector
func NewDetector(asteroidMap Map) *Detector {
	return &Detector{
		asteroidMap: asteroidMap,
	}
}

// FindBestLocation find the best location to build an asteroid detection station
func (d *Detector) FindBestLocation() (math.Point, int) {
	bestLocation := math.NewPoint(-1, -1)
	maxCount := -1
	asteroidCount := d.CountVisibleAsteroids()
	for location, count := range asteroidCount {
		if maxCount <= count {
			maxCount = count
			bestLocation = location
		}
	}
	return bestLocation, maxCount
}

// CountVisibleAsteroids returns a count of visible asteroids for each position on the map
func (d *Detector) CountVisibleAsteroids() map[math.Point]int {
	result := map[math.Point]int{}
	asteroids := d.asteroidMap.Asteroids()
	for _, location := range asteroids {
		result[location] = d.countVisibleAsteroidsFromLocation(location, asteroids)
	}
	return result
}

func (d *Detector) countVisibleAsteroidsFromLocation(location math.Point, asteroids []math.Point) int {
	slopes := map[slope]struct{}{}
	for _, asteroid := range asteroids {
		if asteroid == location {
			continue
		}
		slope := calculateSlope(asteroid, location)
		if _, ok := slopes[slope]; !ok {
			slopes[slope] = struct{}{}
		}
	}

	return len(slopes)
}
