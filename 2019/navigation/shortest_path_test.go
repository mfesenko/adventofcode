package navigation

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_navigation "github.com/mfesenko/adventofcode/.mocks/2019/navigation"
	"github.com/stretchr/testify/assert"
)

func withShortestPathAlg(t *testing.T, test func(*ShortestPathAlg, *mock_navigation.MockGraph)) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	graph := mock_navigation.NewMockGraph(ctrl)
	test(NewShortestPathAlg(graph), graph)
}

func TestShortestPathAlg(t *testing.T) {
	withShortestPathAlg(t, func(s *ShortestPathAlg, graph *mock_navigation.MockGraph) {
		graph.EXPECT().Neighbours("a").Return([]string{"b", "c"})
		graph.EXPECT().Neighbours("b").Return([]string{"a", "d"})
		graph.EXPECT().Neighbours("c").Return([]string{"a"})
		graph.EXPECT().Neighbours("d").Return([]string{"b", "e"})
		graph.EXPECT().Neighbours("e").Return([]string{"d"})

		s.Run("a")

		assert.Equal(t, 1, s.GetDistance("b"))
		assert.Equal(t, 1, s.GetDistance("c"))
		assert.Equal(t, 2, s.GetDistance("d"))
		assert.Equal(t, 3, s.GetDistance("e"))
	})
}
