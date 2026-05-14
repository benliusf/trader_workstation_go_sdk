package send

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilder(t *testing.T) {
	msgId, msgBytes := int32(222), []byte("hello mars")

	expectedLength := []byte{0, 0, 0, 14}
	expectedMsgId := []byte{0, 0, 0, 222}
	expectedMsg := []byte{104, 101, 108, 108, 111, 32, 109, 97, 114, 115}

	bd := newBuilder(len(msgBytes))

	assert.NoError(t, bd.writeMsgId(msgId))
	assert.NoError(t, bd.writeMsgBytes(msgBytes))

	res := bd.bytes()
	assert.Equal(t, len(msgBytes)+8, len(res))
	assert.Equal(t, expectedLength, res[0:4])
	assert.Equal(t, expectedMsgId, res[4:8])
	assert.Equal(t, expectedMsg, res[8:])

	assert.NoError(t, bd.writeMsgId(msgId))
	assert.NoError(t, bd.writeMsgBytes(msgBytes))

	res = bd.bytes()
	assert.Equal(t, len(msgBytes)+8, len(res))
	assert.Equal(t, expectedLength, res[0:4])
	assert.Equal(t, expectedMsgId, res[4:8])
	assert.Equal(t, expectedMsg, res[8:])
}
