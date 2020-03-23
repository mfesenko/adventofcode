package decode

import (
	"errors"
	"strconv"
	"strings"
)

const _offsetSize = 7

// Message represents a message received from communication system
type Message struct {
	data   []int8
	offset int
}

// NewMessage creates a message
func NewMessage(data []int8, offset int) Message {
	return Message{
		data:   data,
		offset: offset,
	}
}

// Data returns message data
func (m Message) Data() []int8 {
	return m.data
}

// Size returns the size of the message
func (m Message) Size() int {
	return len(m.data)
}

// Offset returns offset of the message
func (m Message) Offset() int {
	return m.offset
}

// String returns string representation of the message
func (m Message) String() string {
	builder := &strings.Builder{}
	for _, d := range m.data {
		builder.WriteString(strconv.FormatInt(int64(d), 10))
	}
	return builder.String()
}

// ParseMessage parses a message from a string
func ParseMessage(input string) (Message, error) {
	size := len(input)
	if size < _offsetSize {
		return Message{}, errors.New("message is too short")
	}

	offset, err := strconv.ParseInt(input[:_offsetSize], 10, 64)
	if err != nil {
		return Message{}, err
	}

	data := make([]int8, size)
	for i, r := range input {
		value, err := strconv.ParseInt(string(r), 10, 64)
		if err != nil {
			return Message{}, err
		}

		data[i] = int8(value)
	}

	return NewMessage(data, int(offset)), nil
}
