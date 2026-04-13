package read

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessage(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	buf.Write([]byte("hello mars\000"))
	binary.Write(buf, binary.BigEndian, int32(222))
	buf.Write([]byte("\000"))
	buf.Write([]byte("from earth"))

	msg := &Message{
		body: buf.Bytes(),
	}

	val1, err := msg.ReadStr()
	assert.NoError(t, err)
	assert.Equal(t, "hello mars", val1)
	assert.Equal(t, 11, msg.idx)

	val2, err := msg.ReadMsgId()
	assert.NoError(t, err)
	assert.Equal(t, int32(222)-PROTOBUF_MSG_ID, val2)
	assert.Equal(t, 16, msg.idx)

	val3, err := msg.ReadStr()
	assert.NoError(t, err)
	assert.Equal(t, "from earth", val3)
	assert.Equal(t, 26, msg.idx)

	val4, err := msg.ReadMsgId()
	assert.NoError(t, err)
	assert.Equal(t, int32(222)-PROTOBUF_MSG_ID, val4)
}

func TestReadInt32FromStr(t *testing.T) {
	tests := []struct {
		body     []byte
		expected int32
	}{
		{
			body:     []byte("222"),
			expected: 222,
		},
		{
			body:     []byte("222\000123"),
			expected: 222,
		},
		{
			body:     []byte(fmt.Sprintf("%d", math.MaxInt32)),
			expected: math.MaxInt32,
		},
	}
	for _, tt := range tests {
		msg := &Message{
			body: tt.body,
		}
		actual, err := msg.ReadInt32FromStr()
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, tt.expected, actual)
	}
}
