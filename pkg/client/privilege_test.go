package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermissions(t *testing.T) {
	var ops Permissions
	assert.False(t, CanRead(ops))
	assert.False(t, CanCreate(ops))
	assert.False(t, CanDelete(ops))

	ops = Read
	assert.True(t, CanRead(ops))
	assert.False(t, CanCreate(ops))
	assert.False(t, CanDelete(ops))

	ops = Read | Create
	assert.True(t, CanRead(ops))
	assert.True(t, CanCreate(ops))
	assert.False(t, CanDelete(ops))

	ops = Read | Create | Delete
	assert.True(t, CanRead(ops))
	assert.True(t, CanCreate(ops))
	assert.True(t, CanDelete(ops))
}

func TestRole(t *testing.T) {
	none := Role{}
	assert.True(t, none.None())

	readOnly := ReadOnly()
	assert.True(t, CanRead(readOnly.Orders))
	assert.False(t, CanCreate(readOnly.Orders))
	assert.False(t, CanDelete(readOnly.Orders))
}
