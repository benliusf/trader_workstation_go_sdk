package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestState(t *testing.T) {
	st := &clientState{}

	assert.False(t, st.isReady())
	assert.Equal(t, int32(-1), st.getNextReqId())
	assert.Equal(t, int32(-1), st.getNextOrderId())

	st.setReady()

	assert.True(t, st.isReady())
	assert.Equal(t, int32(0), st.getNextReqId())
	assert.Equal(t, int32(0), st.getNextOrderId())

	st.setNextOrderId(100)

	assert.Equal(t, int32(1), st.getNextReqId())
	assert.Equal(t, int32(100), st.getNextOrderId())
}
