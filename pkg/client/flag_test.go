package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermissions(t *testing.T) {
	var ops Permissions
	assert.False(t, CanRead(ops))
	assert.False(t, CanWrite(ops))

	ops = Read
	assert.True(t, CanRead(ops))
	assert.False(t, CanWrite(ops))

	ops = Read | Write
	assert.True(t, CanRead(ops))
	assert.True(t, CanWrite(ops))
}
