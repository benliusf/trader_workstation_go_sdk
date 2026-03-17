package read

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadStr(t *testing.T) {
	tests := []struct {
		body     []byte
		start    int
		length   int
		expected string
	}{
		{
			body:     []byte("hello world"),
			start:    0,
			length:   5,
			expected: "hello",
		},
		{
			body:     []byte("hello\000world"),
			start:    0,
			length:   -1,
			expected: "hello",
		},
		{
			body:     []byte("hello world"),
			start:    0,
			length:   -1,
			expected: "hello world",
		},
	}
	for _, tt := range tests {
		actual, _ := readStr(tt.body, tt.start, tt.length)
		assert.Equal(t, tt.expected, actual)
	}
}

func TestReadInt32(t *testing.T) {
	tests := []struct {
		body     []byte
		start    int
		expected int32
	}{
		{
			body:     []byte{0, 0, 0, 222},
			start:    0,
			expected: 222,
		},
		{
			body:     []byte{1, 255, 1, 255},
			start:    0,
			expected: 33489407,
		},
		{
			body:     []byte{1, 255, 1, 255, 0, 0, 0, 222},
			start:    4,
			expected: 222,
		},
	}
	for _, tt := range tests {
		actual, _ := readInt32(tt.body, tt.start)
		assert.Equal(t, tt.expected, actual)
	}
}
