package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPaperTrading(t *testing.T) {
	assert.False(t, IsPaperTrading(""))
	assert.False(t, IsPaperTrading("123"))
	assert.True(t, IsPaperTrading("DU123"))
}
