package ascii

import (
	"github.com/mfesenko/adventofcode/2019/math"
)

const (
	scaffold = '#'
	empty    = '.'
	newLine  = '\n'
)

// MapLoader is responsible for loading a ScaffoldMap from runes
type MapLoader struct {
	scaffoldMap *ScaffoldMap
	done        chan bool
	x           int64
	y           int64
}

// NewMapLoader creates a new MapLoader
func NewMapLoader() *MapLoader {
	return &MapLoader{
		scaffoldMap: NewScaffoldMap(),
		done:        make(chan bool),
	}
}

// Wait waits until the loader is done
func (l *MapLoader) Wait() {
	<-l.done
}

// Done returns a boolean flag indicating if the scaffold map was loaded
func (l *MapLoader) Done() bool {
	select {
	case <-l.done:
		return true
	default:
		return false
	}
}

// ScaffoldMap return a map
func (l *MapLoader) ScaffoldMap() *ScaffoldMap {
	return l.scaffoldMap
}

// ProcessRune updates a map
func (l *MapLoader) ProcessRune(r rune) {
	if l.Done() {
		return
	}

	switch r {
	case newLine:
		if l.x == 0 {
			close(l.done)
			return
		}
		l.y++
		l.x = 0

	case empty:
		l.x++

	case scaffold:
		l.scaffoldMap.AddScaffold(math.NewPoint(l.x, l.y))
		l.x++

	default:
		position := math.NewPoint(l.x, l.y)
		l.scaffoldMap.AddScaffold(position)
		l.scaffoldMap.SetDirection(DirectionFromRune(r))
		l.scaffoldMap.SetPosition(position)
		l.x++
	}
}
