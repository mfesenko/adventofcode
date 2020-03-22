package repair

import (
	"sync"

	"github.com/mfesenko/adventofcode/2019/async"
	"github.com/mfesenko/adventofcode/2019/math"
	"github.com/mfesenko/adventofcode/2019/navigation"
)

type (
	// Computer represents a computer that provides instructions to the Droid
	Computer interface {
		Execute()
		Input() chan int64
		Output() chan int64
		Stop()
	}

	// Droid represents a repair droid
	Droid struct {
		computer             Computer
		position             math.Point
		oxygenSystemPosition math.Point
		areaMap              *areaMap
		moves                []MovementCommand
	}
)

// NewDroid creates a new Droid
func NewDroid(computer Computer) *Droid {
	return &Droid{
		computer:             computer,
		position:             math.NewPoint(0, 0),
		oxygenSystemPosition: math.NewPoint(0, 0),
		areaMap:              newAreaMap(),
	}
}

// Explore sends a repair droid to explore the area around it
func (d *Droid) Explore() {
	executor := async.NewExecutor(d.computer)
	done := &sync.WaitGroup{}
	done.Add(1)
	executor.ExecuteAsync(done)

	d.areaMap.Update(d.position, MovedOneStep)
	for {
		command := d.nextCommand()
		if command == Invalid {
			break
		}
		d.move(command)
	}

	d.computer.Stop()
	done.Wait()
}

func (d *Droid) nextCommand() MovementCommand {
	moves := []MovementCommand{
		North, East, South, West,
	}
	for _, command := range moves {
		point := command.Apply(d.position)
		if !d.areaMap.Contains(point) {
			d.moves = append(d.moves, command)
			return command
		}
	}

	if len(d.moves) == 0 {
		return Invalid
	}

	index := len(d.moves) - 1
	lastMove := d.moves[index]
	d.moves = d.moves[:index]
	return lastMove.Reverse()
}

func (d *Droid) move(command MovementCommand) {
	d.computer.Input() <- int64(command)
	nextPosition := command.Apply(d.position)
	statusCode := StatusCode(<-d.computer.Output())
	d.areaMap.Update(nextPosition, statusCode)
	d.processStatusCode(nextPosition, statusCode)
}

func (d *Droid) processStatusCode(position math.Point, statusCode StatusCode) {
	switch statusCode {
	case HitWall:
		d.moves = d.moves[:len(d.moves)-1]

	case MovedOneStep:
		d.position = position

	case FoundOxygenSystem:
		d.position = position
		d.oxygenSystemPosition = position
	}
}

// ShortestPathToOxygenSystem returns the minimal amount of steps required to move
// the droid from starting position to the oxygen system
func (d *Droid) ShortestPathToOxygenSystem() int {
	s := navigation.NewShortestPathAlg(d.areaMap)
	s.Run(math.NewPoint(0, 0).String())
	return s.GetDistance(d.oxygenSystemPosition.String())
}

// MinutesToFillWithOxygen returns the amount of minutes it will take to fill the area with oxygen
func (d *Droid) MinutesToFillWithOxygen() int {
	s := navigation.NewShortestPathAlg(d.areaMap)
	s.Run(d.oxygenSystemPosition.String())
	max := -1
	for key, tile := range d.areaMap.data {
		if tile.statusCode == HitWall {
			continue
		}
		distance := s.GetDistance(key)
		if max < distance {
			max = distance
		}
	}
	return max
}
