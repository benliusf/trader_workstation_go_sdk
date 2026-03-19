package read

import (
	"fmt"
	"io"
	"strconv"
	"sync"
)

type Message struct {
	mu   sync.Mutex
	body []byte
	idx  int
	id   int32
}

func MessageFromBytes(b []byte) (*Message, error) {
	if len(b) == 0 {
		return nil, fmt.Errorf("empty bytes")
	}
	return &Message{
		body: b,
		id:   -1,
	}, nil
}

func (m *Message) updateIndex(incr int) {
	m.idx += incr
	if m.idx < len(m.body)-1 &&
		m.body[m.idx] == 0 {
		m.idx++
	}
}

func (m *Message) ReadBytes() ([]byte, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	b := m.body[m.idx:]
	m.updateIndex(len(b))
	return b, nil
}

func (m *Message) ReadMsgId() (int32, error) {
	if m.id > 0 {
		return m.id, nil
	}
	raw, err := m.ReadInt32()
	if err != nil {
		return -1, err
	}
	m.id = raw - PROTOBUF_MSG_ID
	return m.id, nil
}

func (m *Message) ReadStr() (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.idx >= len(m.body) {
		return "", io.EOF
	}
	val, err := readStr(m.body, m.idx, -1)
	if err != nil {
		return "", err
	}
	m.updateIndex(len(val))
	return val, nil
}

func (m *Message) ReadInt32() (int32, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.idx >= len(m.body) {
		return -1, io.EOF
	}
	val, err := readInt32(m.body, m.idx)
	if err != nil {
		return -1, err
	}
	m.updateIndex(int32BytesLength)
	return val, nil
}

func (m *Message) ReadInt32FromStr() (int32, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.idx >= len(m.body) {
		return -1, io.EOF
	}
	tmpStr, err := readStr(m.body, m.idx, -1)
	if err != nil {
		return -1, err
	}
	tmpInt, err := strconv.ParseInt(tmpStr, 10, 32)
	if err != nil {
		return -1, err
	}
	m.updateIndex(len(tmpStr))
	return int32(tmpInt), nil
}
