package ascii

import (
	"fmt"
	"sync"

	"github.com/mfesenko/adventofcode/2019/async"
	"github.com/mfesenko/adventofcode/2019/math"
)

type (
	// Computer represents a computer that provides current view of the scaffolds to the robot
	Computer interface {
		Execute()
		Input() chan int64
		Output() chan int64
	}

	// Robot represents a vacuum robot
	Robot struct {
		computer       Computer
		routineBuilder *MainRoutineBuilder
	}
)

// NewRobot creates a new robot
func NewRobot(computer Computer) *Robot {
	return &Robot{
		computer:       computer,
		routineBuilder: NewMainRoutineBuilder(),
	}
}

// Run starts a robot
func (r *Robot) Run(withVideo bool) (int64, int64) {
	executor := async.NewExecutor(r.computer)
	done := &sync.WaitGroup{}
	done.Add(1)
	executor.ExecuteAsync(done)
	defer done.Wait()

	mapLoader := NewMapLoader()
	result := make(chan int64)

	go r.processOutput(mapLoader, result)

	mapLoader.Wait()
	scaffoldMap := mapLoader.ScaffoldMap()
	alignmentParameters := r.alignmentParameters(scaffoldMap.Intersections())
	commands := r.buildInputCommands(scaffoldMap, withVideo)
	r.sendInput(commands)

	return alignmentParameters, <-result
}

func (r *Robot) processOutput(mapLoader *MapLoader, result chan int64) {
	for {
		v, ok := <-r.computer.Output()
		if !ok {
			break
		}

		if !isASCII(v) {
			result <- v
			break
		}

		r := rune(v)
		fmt.Printf("%c", r)

		mapLoader.ProcessRune(r)
	}

	close(result)
}

func (r *Robot) alignmentParameters(intersections []math.Point) int64 {
	result := int64(0)
	for _, i := range intersections {
		result += i.X * i.Y
	}
	return result
}

func (r *Robot) buildInputCommands(scaffoldMap *ScaffoldMap, withVideo bool) []string {
	path := scaffoldMap.GeneratePath()
	r.routineBuilder.Build(path, []string{"A", "B", "C"})
	input := []string{
		r.routineBuilder.MainRoutine(),
		r.routineBuilder.Function("A"),
		r.routineBuilder.Function("B"),
		r.routineBuilder.Function("C"),
	}
	videoCommand := "n"
	if withVideo {
		videoCommand = "y"
	}
	input = append(input, videoCommand)
	return input
}

func (r *Robot) sendInput(commands []string) {
	for _, command := range commands {
		for _, c := range command {
			r.computer.Input() <- int64(c)
		}
		r.computer.Input() <- newLine
	}
}
