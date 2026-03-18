package net

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

const (
	API_HEADER = "API\000"

	MIN_VER = MIN_SERVER_VER_FRACTIONAL_LAST_SIZE
	MAX_VER = MIN_SERVER_VER_FRACTIONAL_LAST_SIZE
)

var verHeader string

func init() {
	verHeader = fmt.Sprintf("%v..%v", MIN_VER, MAX_VER)
}

type TWSConn struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	conn net.Conn
}

func (c *TWSConn) Open(addr string) (err error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	c.conn = conn
	defer func() {
		if err != nil {
			c.Close()
		}
	}()
	if err := c.sendConnectRequest(); err != nil {
		return err
	}
	return nil
}

func (c *TWSConn) sendConnectRequest() error {
	buf := bytes.NewBuffer(nil)
	if _, err := buf.Write([]byte(API_HEADER)); err != nil {
		return err
	}
	ver := []byte(verHeader)
	if err := binary.Write(buf, binary.BigEndian, int32(len(ver))); err != nil {
		return err
	}
	if _, err := buf.Write(ver); err != nil {
		return err
	}
	if _, err := c.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}

func (c *TWSConn) Close() error {
	if c != nil &&
		c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func (c *TWSConn) Read(buf []byte) (int, error) {
	if c.ReadTimeout > 0 {
		deadline := time.Now().Add(c.ReadTimeout)
		c.conn.SetReadDeadline(deadline)
	}
	return c.conn.Read(buf)
}

func (c *TWSConn) Write(b []byte) (int, error) {
	if c.WriteTimeout > 0 {
		deadline := time.Now().Add(c.WriteTimeout)
		c.conn.SetWriteDeadline(deadline)
	}
	return c.conn.Write(b)
}
