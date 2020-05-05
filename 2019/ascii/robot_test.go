package ascii

import (
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	mock_ascii "github.com/mfesenko/adventofcode/.mocks/2019/ascii"
	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

type mocks struct {
	computer *mock_ascii.MockComputer
	input    chan int64
	output   chan int64
}

func withRobot(t *testing.T, test func(*Robot, mocks)) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	defer goleak.VerifyNone(t)

	output := make(chan int64, 1)
	input := make(chan int64, 1000)
	computer := mock_ascii.NewMockComputer(ctrl)
	computer.EXPECT().Output().AnyTimes().Return(output)
	computer.EXPECT().Input().AnyTimes().Return(input)
	test(NewRobot(computer), mocks{
		computer: computer,
		input:    input,
		output:   output,
	})
}

func TestRobot(t *testing.T) {
	withRobot(t, func(robot *Robot, m mocks) {
		testDust := int64(789)
		testParameters := int64(276)
		m.computer.EXPECT().Execute().Do(func() {
			for _, r := range testComputerOutput() {
				m.output <- int64(r)
			}

			m.output <- testDust
		})

		parameters, dust := robot.Run(true)

		assert.Equal(t, testParameters, parameters)
		assert.Equal(t, testDust, dust)

		input := readLinesFromChannel(m.input, 5)
		assert.Equal(t, []string{
			"A,B,C,B,A,C",
			"R,8,R,8",
			"R,4,R,4",
			"R,8,L,6,L,2",
			"y",
		}, input)

	})

}

func testComputerOutput() string {
	return `#######...#####
#.....#...#...#
#.....#...#...#
......#...#...#
......#...###.#
......#.....#.#
^########...#.#
......#.#...#.#
......#########
........#...#..
....#########..
....#...#......
....#...#......
....#...#......
....#####......

Hello world

#######...#####
#.....#...#...#
v.....#...#...#
......#...#...#
......#...###.#
......#.....#.#
#########...#.#
......#.#...#.#
......#########
........#...#..
....#########..
....#...#......
....#...#......
....#...#......
....#####......

`
}

func readLinesFromChannel(channel chan int64, expectedCount int) []string {
	lines := make([]string, 0)
	inputBuilder := &strings.Builder{}
	count := 0
	for count < expectedCount {
		v, ok := <-channel
		if !ok {
			break
		}
		r := rune(v)
		if r == newLine {
			lines = append(lines, inputBuilder.String())
			count++
			inputBuilder = &strings.Builder{}
			continue
		}
		inputBuilder.WriteRune(r)
	}
	return lines
}
