package drawing

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/golang/mock/gomock"
	mock_drawing "github.com/mfesenko/adventofcode/.mocks/2019/drawing"
	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

type (
	mocks struct {
		computer *mock_drawing.MockComputer
		input    chan int64
		output   chan int64
	}

	testCommand struct {
		input   int64
		color   int64
		command int64
	}
)

func withRobot(t *testing.T, test func(*Robot, mocks)) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	defer goleak.VerifyNone(t)

	input := make(chan int64, 20)
	output := make(chan int64, 20)
	computer := mock_drawing.NewMockComputer(ctrl)
	robot := NewRobot(computer)
	test(robot, mocks{
		computer: computer,
		input:    input,
		output:   output,
	})
}

func TestChangeColorAddsATileWhenCalledOnUncoveredTile(t *testing.T) {
	withRobot(t, func(robot *Robot, m mocks) {
		assert.Equal(t, 0, len(robot.covered))
		assert.Equal(t, Black, robot.currentColor())

		robot.ChangeColor(White)

		assert.Equal(t, 1, len(robot.covered))
		tile := robot.covered[robot.position]
		assert.Equal(t, White, tile.color)
		assert.Equal(t, 1, tile.count)
	})
}

func TestChangeColorUpdatesATileWhenCalledOnCoveredTile(t *testing.T) {
	withRobot(t, func(robot *Robot, m mocks) {
		count := rand.Int()
		robot.covered[robot.position] = &tile{
			color: White,
			count: count,
		}
		assert.Equal(t, White, robot.currentColor())

		robot.ChangeColor(Black)

		assert.Equal(t, 1, len(robot.covered))
		tile := robot.covered[robot.position]
		assert.Equal(t, Black, tile.color)
		assert.Equal(t, count+1, tile.count)
	})
}

func TestRobotPanicsOnUnsupportedDirectionCommand(t *testing.T) {
	withRobot(t, func(robot *Robot, m mocks) {
		command := TurnRight + 1
		m.computer.EXPECT().Input().AnyTimes().Return(m.input)
		m.computer.EXPECT().Output().AnyTimes().Return(m.output)
		m.computer.EXPECT().Execute().Do(func() {
			m.output <- int64(White)
			m.output <- int64(command)
		})
		assert.PanicsWithValue(t, fmt.Sprintf(unsupportedDirectionCommand, command), func() {
			robot.Run()
		})
	})
}

func TestRobotRun(t *testing.T) {
	withRobot(t, func(robot *Robot, m mocks) {
		m.computer.EXPECT().Input().AnyTimes().Return(m.input)
		m.computer.EXPECT().Output().AnyTimes().Return(m.output)
		tests := testsCommands()
		m.computer.EXPECT().Execute().Do(func() {
			for _, test := range tests {
				input := <-m.input
				assert.Equal(t, test.input, input)
				m.output <- test.color
				m.output <- test.command
			}
		})

		robot.Run()

		assert.Equal(t, 6, robot.CoveredTiles())
		expectedImage := [][]uint16{
			{White, Black, Black},
			{White, Black, Black},
			{Black, White, White},
		}
		assert.Equal(t, expectedImage, robot.Image().Render())
	})
}

func testsCommands() []testCommand {
	return []testCommand{
		{
			input:   0,
			color:   1,
			command: 0,
		},
		{
			input:   0,
			color:   0,
			command: 0,
		},
		{
			input:   0,
			color:   1,
			command: 0,
		},
		{
			input:   0,
			color:   1,
			command: 0,
		},
		{
			input:   1,
			color:   0,
			command: 1,
		},
		{
			input:   0,
			color:   1,
			command: 0,
		},
		{
			input:   0,
			color:   1,
			command: 0,
		},
	}
}
