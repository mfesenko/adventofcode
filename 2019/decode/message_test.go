package decode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenInputIsTooShortThenParseMessageReturnsAnError(t *testing.T) {
	_, err := ParseMessage("1234")

	assert.Error(t, err)
}

func TestWhenUnableToParseOffsetThenParseMessageReturnsAnError(t *testing.T) {
	_, err := ParseMessage("a1234567")

	assert.Error(t, err)
}

func TestWhenUnableToParseMessageDataThenParseMessageReturnsAnError(t *testing.T) {
	_, err := ParseMessage("12345678a")

	assert.Error(t, err)
}

func TestWhenInputIsValidThenParseMessageReturnsAMessage(t *testing.T) {
	message, err := ParseMessage("0123456789")

	assert.NoError(t, err)
	assert.Equal(t, 123456, message.Offset())
	assert.Equal(t, 10, message.Size())
	assert.Equal(t, []int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, message.Data())
}
