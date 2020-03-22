package repair

import (
	"github.com/mfesenko/adventofcode/2019/math"
)

type (
	tile struct {
		position   math.Point
		statusCode StatusCode
	}

	areaMap struct {
		data map[string]*tile
	}
)

func newTile(position math.Point, statusCode StatusCode) *tile {
	return &tile{
		position:   position,
		statusCode: statusCode,
	}
}

func (t *tile) Neighbours() []math.Point {
	return []math.Point{
		math.NewPoint(t.position.X, t.position.Y+1),
		math.NewPoint(t.position.X, t.position.Y-1),
		math.NewPoint(t.position.X+1, t.position.Y),
		math.NewPoint(t.position.X-1, t.position.Y),
	}
}

func newAreaMap() *areaMap {
	return &areaMap{
		data: map[string]*tile{},
	}
}

func (m *areaMap) Update(position math.Point, statusCode StatusCode) {
	m.data[position.String()] = newTile(position, statusCode)
}

func (m *areaMap) Contains(position math.Point) bool {
	_, ok := m.data[position.String()]
	return ok
}

func (m *areaMap) Neighbours(nodeName string) []string {
	tile, ok := m.data[nodeName]
	if !ok {
		return nil
	}

	var neighbours []string
	for _, n := range tile.Neighbours() {
		if tile, ok := m.data[n.String()]; ok && tile.statusCode != HitWall {
			neighbours = append(neighbours, n.String())
		}
	}
	return neighbours
}
