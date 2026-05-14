package send

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt32Bytes(t *testing.T) {
	assert.Equal(t, []byte{0, 0, 0, 201}, int32Bytes(201))
	assert.Equal(t, []byte{0, 0, 1, 0}, int32Bytes(256))
	assert.Equal(t, []byte{127, 255, 255, 255}, int32Bytes(math.MaxInt32))
}
