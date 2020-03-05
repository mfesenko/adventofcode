package math

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxReturnsGreaterValue(t *testing.T) {
	a := rand.Int63()
	b := a + 1

	assert.Equal(t, b, Max(a, b))
	assert.Equal(t, b, Max(b, a))
}

func TestMinReturnsSmallerValue(t *testing.T) {
	a := rand.Int63()
	b := a + 1

	assert.Equal(t, a, Min(a, b))
	assert.Equal(t, a, Min(b, a))
}

func TestAbsReturnsAbsoluteValue(t *testing.T) {
	a := rand.Int63()

	assert.Equal(t, a, Abs(a))
	assert.Equal(t, a, Abs(-a))
}

func TestBetweenReturnsTrueIfTheThirdArgumentIsBetweenTheFirstTwo(t *testing.T) {
	a := rand.Int63()
	b := a + 1
	c := b + 1

	assert.True(t, Between(a, c, b))
	assert.True(t, Between(c, a, b))
	assert.True(t, Between(a, a, a))
	assert.False(t, Between(a, b, c))
}

func TestGCD(t *testing.T) {
	assert.Equal(t, int64(22), GCD(0, 22))
	assert.Equal(t, int64(22), GCD(22, 0))
	assert.Equal(t, int64(0), GCD(0, 0))
	assert.Equal(t, int64(1), GCD(13, 7))
	assert.Equal(t, int64(7), GCD(21, 7))
	assert.Equal(t, int64(3), GCD(21, 18))
}

func TestLCM(t *testing.T) {
	assert.Equal(t, int64(15), LCM(1, 15))
	assert.Equal(t, int64(30), LCM(2, 15))
	assert.Equal(t, int64(18), LCM(6, 9))
	assert.Equal(t, int64(0), LCM(0, 9))
}
