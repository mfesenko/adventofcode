package chemistry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUpdatesAmontOfChemicalInAFlask(t *testing.T) {
	chemical := "asdf"
	flask := newFlask()
	assert.NotContains(t, flask, chemical)

	amount := int64(15)
	flask.Add(chemical, amount)
	assert.Equal(t, amount, flask[chemical])

	increment := int64(3)
	flask.Add(chemical, increment)
	assert.Equal(t, amount+increment, flask[chemical])
}

func TestContainsOnly(t *testing.T) {
	chemical := "asdf"
	otherChemical := "qwerty"
	flask := newFlask()
	assert.True(t, flask.ContainsOnly(chemical))

	flask.Add(chemical, 22)
	assert.True(t, flask.ContainsOnly(chemical))
	assert.False(t, flask.ContainsOnly(otherChemical))

	flask.Add(otherChemical, 12)
	assert.False(t, flask.ContainsOnly(chemical))
}
