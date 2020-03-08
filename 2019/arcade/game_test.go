package arcade

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_arcade "github.com/mfesenko/adventofcode/.mocks/2019/arcade"
	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

type mocks struct {
	computer *mock_arcade.MockComputer
	output   chan int64
	input    chan int64
}

func withGame(t *testing.T, test func(*Game, mocks)) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	defer goleak.VerifyNone(t)

	computer := mock_arcade.NewMockComputer(ctrl)
	output := make(chan int64, 1)
	input := make(chan int64, 1)

	game := NewGame(computer)
	test(game, mocks{
		computer: computer,
		output:   output,
		input:    input,
	})
}

func TestCountTiles(t *testing.T) {
	withGame(t, func(game *Game, m mocks) {
		assert.Equal(t, 0, game.CountTiles(BlockTile))

		expectExecuteWithOutput(m, []int64{
			1, 2, int64(HorizontalPaddleTile),
			6, 5, int64(BlockTile),
			1, 2, int64(BlockTile),
			4, 6, int64(WallTile),
		})
		game.Play()

		assert.Equal(t, 2, game.CountTiles(BlockTile))
	})
}

func TestUpdateScoreInstructionUpdatesTheGameScore(t *testing.T) {
	withGame(t, func(game *Game, m mocks) {
		assert.Equal(t, int64(0), game.Score())

		score := int64(345789)
		expectExecuteWithOutput(m, []int64{-1, 0, score})

		game.Play()

		assert.Equal(t, score, game.Score())
	})
}

func TestWhenBallPositionIsUpdatedThenJoystickSignalIsSent(t *testing.T) {
	type testData struct {
		output []int64
		input  int64
	}
	tests := []testData{
		{
			output: []int64{
				3, 3, int64(HorizontalPaddleTile),
				0, 0, int64(BallTile),
			},
			input: int64(JoystickLeft),
		},
		{
			output: []int64{
				3, 3, int64(HorizontalPaddleTile),
				3, 1, int64(BallTile),
			},
			input: int64(JoystickNeutral),
		},
		{
			output: []int64{
				3, 3, int64(HorizontalPaddleTile),
				6, 2, int64(BallTile),
			},
			input: int64(JoystickRight),
		},
	}
	for _, test := range tests {
		withGame(t, func(game *Game, m mocks) {
			expectExecuteWithOutput(m, test.output)
			m.computer.EXPECT().Input().Return(m.input)

			game.Play()

			assert.Equal(t, 1, len(m.input))
			assert.Equal(t, test.input, <-m.input)
		})
	}
}

func expectExecuteWithOutput(m mocks, output []int64) {
	m.computer.EXPECT().Output().AnyTimes().Return(m.output)
	m.computer.EXPECT().Execute().Do(func() {
		for _, o := range output {
			m.output <- o
		}
		close(m.output)
	})
}
