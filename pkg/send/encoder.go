package send

import "encoding/binary"

const int32BytesLength int = 4

func int32Bytes(v int32) []byte {
	tmp := make([]byte, int32BytesLength)
	binary.BigEndian.PutUint32(tmp, uint32(v))
	return tmp[:]
}
