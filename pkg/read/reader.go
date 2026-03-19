package read

import (
	"encoding/binary"
	"time"

	"github.com/benliusf/trader_workstation_go_sdk/pkg/net"
)

func Next(conn *net.TWSConn) (*Message, error) {
	buf := make([]byte, 4)
	_, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}
	tempInt := binary.BigEndian.Uint32(buf)
	lengthHeader := int32(tempInt)
	buf = make([]byte, lengthHeader)
	_, err = conn.Read(buf)
	if err != nil {
		return nil, err
	}
	return MessageFromBytes(buf)
}

func ServerVersion(conn *net.TWSConn) (*ServerVersionResponse, error) {
	const tsFormat = "20060102 15:04:05 PST"
	msg, err := Next(conn)
	if err != nil {
		return nil, err
	}
	serverVersion, err := msg.ReadInt32FromStr()
	if err != nil {
		return nil, err
	}
	var serverTs int64
	tsStr, _ := msg.ReadStr()
	ts, _ := time.Parse(tsFormat, tsStr)
	if !ts.IsZero() {
		serverTs = ts.Unix()
	}
	return &ServerVersionResponse{
		ServerVersion: serverVersion,
		Timestamp:     serverTs,
	}, nil
}
