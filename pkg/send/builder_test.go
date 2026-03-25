package send

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilder(t *testing.T) {
	msgId := int32(222)
	msg := []byte("hello mars")
	bd := newBuilder().withMsgId(msgId).withMsgBytes(msg)
	actual, err := bd.build()
	assert.Nil(t, err)
	assert.Equal(t, []byte{0, 0, 0, 14}, actual[0:4])                                   // assert msg length
	assert.Equal(t, []byte{0, 0, 0, 222}, actual[4:8])                                  // assert msgId
	assert.Equal(t, []byte{104, 101, 108, 108, 111, 32, 109, 97, 114, 115}, actual[8:]) // assert msg bytes
}
