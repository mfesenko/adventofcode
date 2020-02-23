package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenGivenNilSliceThenCopyInt64ReturnsNil(t *testing.T) {
	assert.Nil(t, CopyInt64(nil))
}

func TestWhenGivenNotNilSliceThenCopyInt64ReturnsACopy(t *testing.T) {
	testData := []int64{1, 2, 3}

	copy := CopyInt64(testData)

	assert.Equal(t, testData, copy)

	copy[0] = 99
	assert.Equal(t, int64(1), testData[0])
	assert.Equal(t, int64(99), copy[0])
}
