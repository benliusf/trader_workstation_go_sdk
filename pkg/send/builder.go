package send

import (
	"fmt"
)

type builder struct {
	buf []byte

	msgId    int32
	msgBytes []byte
}

func newBuilder(msgSize int) *builder {
	size := msgSize + 8
	return &builder{
		buf: make([]byte, size, size),
	}
}

func (b *builder) bytes() []byte {
	return b.buf
}

func (b *builder) writeMsgId(v int32) error {
	if len(b.buf) < 8 {
		return fmt.Errorf("bad buffer size")
	}
	tmp := int32Bytes(v)
	copy(b.buf[int32BytesLength:8], tmp)
	b.msgId = v
	return nil
}

func (b *builder) writeMsgLength(v int32) error {
	if len(b.buf) < 4 {
		return fmt.Errorf("bad buffer size")
	}
	tmp := int32Bytes(v)
	copy(b.buf[0:4], tmp)
	return nil
}

func (b *builder) writeMsgBytes(v []byte) error {
	if len(b.buf) < len(v)+8 {
		return fmt.Errorf("bad buffer size")
	}
	if err := b.writeMsgLength(int32(int32BytesLength + len(v))); err != nil {
		return err
	}
	copy(b.buf[8:], v)
	b.msgBytes = v
	return nil
}
