package send

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilder(t *testing.T) {
	tests := []struct {
		msgId             int32
		msg               []byte
		expectedMsgLength []byte
		expectedMsgId     []byte
		expectedMsgBytes  []byte
	}{
		{
			msgId:             222,
			msg:               []byte("hello mars"),
			expectedMsgLength: []byte{0, 0, 0, 14},
			expectedMsgId:     []byte{0, 0, 0, 222},
			expectedMsgBytes:  []byte{104, 101, 108, 108, 111, 32, 109, 97, 114, 115},
		},
	}
	for _, tt := range tests {
		bd := newBuilder().withMsgId(tt.msgId).withMsgBytes(tt.msg)
		actual, err := bd.build()
		assert.Nil(t, err)
		assert.Equal(t, []byte{0, 0, 0, 14}, actual[0:4])                                   // assert msg length
		assert.Equal(t, []byte{0, 0, 0, 222}, actual[4:8])                                  // assert msgId
		assert.Equal(t, []byte{104, 101, 108, 108, 111, 32, 109, 97, 114, 115}, actual[8:]) // assert msg bytes
	}
}
