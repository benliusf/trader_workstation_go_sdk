package net

import (
	"net"
	"time"
)

type MockConn struct{}

func NewMockConn() *TWSConn {
	return &TWSConn{
		conn: &MockConn{},
	}
}

func (c *MockConn) Read(b []byte) (n int, err error) {
	return -1, nil
}

func (c *MockConn) Write(b []byte) (n int, err error) {
	return -1, nil
}

func (c *MockConn) Close() error {
	return nil
}

func (c *MockConn) LocalAddr() net.Addr {
	return nil
}

func (c *MockConn) RemoteAddr() net.Addr {
	return nil
}

func (c *MockConn) SetDeadline(t time.Time) error {
	return nil
}

func (c *MockConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (c *MockConn) SetWriteDeadline(t time.Time) error {
	return nil
}
