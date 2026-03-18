package net

import (
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTWSConn(t *testing.T) {
	mockServer, mockClient := net.Pipe()

	twsConn := &TWSConn{
		conn: mockClient,
	}

	verHeader = "123..456"

	expected := []byte(API_HEADER)
	expected = append(expected, []byte{0, 0, 0, 8}...)
	expected = append(expected, []byte(verHeader)...)

	done := make(chan struct{})

	go func() {
		defer func() {
			done <- struct{}{}
		}()
		buf := make([]byte, 1024)
		n, err := mockServer.Read(buf)
		if err != nil {
			t.Error(err)
			return
		}
		assert.Equal(t, expected, buf[:n])
	}()

	if err := twsConn.sendConnectRequest(); err != nil {
		t.Fatal(err)
	}

	select {
	case <-done:
	case <-time.After(10 * time.Second):
		t.Fatal(fmt.Errorf("timed out"))
	}
}
