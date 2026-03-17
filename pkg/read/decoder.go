package read

import (
	"encoding/binary"
	"fmt"
)

func readStr(b []byte, start int, length int) (string, error) {
	var i int = start
	for ; i < len(b); i++ {
		// check for null byte (zero value)
		if b[i] == 0 {
			break
		}
		if length > 0 &&
			(i-start) >= length {
			break
		}
	}
	return string(b[start:i]), nil
}

func readInt32(b []byte, start int) (int32, error) {
	const int32BytesLength int = 4
	bytesLength := int32BytesLength
	if len(b) < int32BytesLength {
		bytesLength = len(b)
	} else if len(b[start:]) < int32BytesLength {
		bytesLength = len(b[start:])
	}
	if bytesLength < 4 {
		return -1, fmt.Errorf("invalid bytes length of '%d' for int32", bytesLength)
	}
	tempInt := binary.BigEndian.Uint32(b[start : start+int32BytesLength])
	return int32(tempInt), nil
}
