package repair

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_repair "github.com/mfesenko/adventofcode/.mocks/2019/repair"
	"github.com/stretchr/testify/assert"
)

func TestDroid(t *testing.T) {
	withComputer(t, func(computer *mock_repair.MockComputer) {
		droid := NewDroid(computer)

		droid.Explore()

		assert.Equal(t, 2, droid.ShortestPathToOxygenSystem())
		assert.Equal(t, 3, droid.MinutesToFillWithOxygen())
	})
}

func withComputer(t *testing.T, test func(*mock_repair.MockComputer)) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Test computer corresponds to the following area map:
	//  ###
	// #.D.#
	// #O##
	//  #
	commands := []MovementCommand{
		North, East, North, East, South, West, South, West, North, South, South, West, North, West, East,
	}
	statusCodes := []StatusCode{
		HitWall, MovedOneStep, HitWall, HitWall, HitWall, MovedOneStep, HitWall, MovedOneStep, HitWall, FoundOxygenSystem, HitWall, HitWall, MovedOneStep, HitWall, MovedOneStep,
	}
	input := make(chan int64)
	output := make(chan int64)
	computer := mock_repair.NewMockComputer(ctrl)
	computer.EXPECT().Input().Return(input).AnyTimes()
	computer.EXPECT().Output().Return(output).AnyTimes()
	computer.EXPECT().Execute().Do(func() {
		for i, command := range commands {
			assert.Equal(t, int64(command), <-input)
			output <- int64(statusCodes[i])
		}
	})
	computer.EXPECT().Stop()
	test(computer)
}
