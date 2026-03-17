package send

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const int32BytesLength int32 = 4

type builder struct {
	buf *bytes.Buffer

	msgId    int32
	msgBytes []byte
}

func newBuilder() *builder {
	return &builder{
		buf: bytes.NewBuffer(nil),
	}
}

func (b *builder) build() ([]byte, error) {
	msgLength := int32BytesLength + int32(len(b.msgBytes))
	if err := binary.Write(b.buf, binary.BigEndian, msgLength); err != nil {
		return nil, fmt.Errorf("failed to write msgLength=%d: %w", msgLength, err)
	}
	if err := binary.Write(b.buf, binary.BigEndian, b.msgId); err != nil {
		return nil, fmt.Errorf("failed to write msgId=%d: %w", b.msgId, err)
	}
	if _, err := b.buf.Write(b.msgBytes); err != nil {
		return nil, fmt.Errorf("failed to write bytes: %w", err)
	}
	return b.buf.Bytes(), nil
}

func (b *builder) withMsgId(v int32) *builder {
	b.msgId = v
	return b
}

func (b *builder) withMsgBytes(v []byte) *builder {
	b.msgBytes = v
	return b
}
