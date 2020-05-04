package intcode

import (
	"strconv"
	"strings"

	"github.com/mfesenko/adventofcode/2019/input"
)

const addressCanNotBeNegative = "address can not be negative"

// Program represents a program that can be executed by Intcode computer
type Program struct {
	code []int64
}

// NewProgram creates a new program
func NewProgram(code []int64) *Program {
	return &Program{
		code: code,
	}
}

// LoadProgram loads program from the file
func LoadProgram(filePath string) (*Program, error) {
	data, err := input.LoadFromFile(filePath)
	if err != nil {
		return nil, err
	}

	words := strings.Split(data[0], ",")
	code := make([]int64, len(words))
	for i, word := range words {
		value, err := strconv.ParseInt(word, 10, 64)
		if err != nil {
			return nil, err
		}
		code[i] = value
	}

	return NewProgram(code), nil
}

// Copy returns a copy of the program
func (p *Program) Copy() *Program {
	code := make([]int64, len(p.code))
	copy(code, p.code)
	return NewProgram(code)
}

// Len returns the length of the program
func (p *Program) Len() int64 {
	return int64(len(p.code))
}

// Read returns a value by address
func (p *Program) Read(address int64) int64 {
	p.validateAddress(address)

	if address < p.Len() {
		return p.code[address]
	}

	return 0
}

// Write writes a value to address
func (p *Program) Write(address int64, value int64) {
	p.validateAddress(address)

	if address >= p.Len() {
		code := make([]int64, address+1)
		copy(code, p.code)
		p.code = code
	}

	p.code[address] = value
}

func (p *Program) validateAddress(address int64) {
	if address < 0 {
		panic(addressCanNotBeNegative)
	}
}

// SetNoun sets the noun for the program
func (p *Program) SetNoun(noun int64) {
	p.Write(1, noun)
}

// SetVerb sets the verb for the program
func (p *Program) SetVerb(verb int64) {
	p.Write(2, verb)
}

// AsString returns string representation of the program
func (p *Program) AsString() string {
	data := make([]string, p.Len())
	for i, value := range p.code {
		data[i] = strconv.FormatInt(value, 10)
	}
	return strings.Join(data, ",")
}
