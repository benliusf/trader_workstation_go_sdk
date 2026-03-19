package client

type Permissions uint64

const (
	Read Permissions = 1 << iota
	Write
)

func CanRead(o Permissions) bool {
	return (o & Read) == Read
}

func CanWrite(o Permissions) bool {
	return (o & Write) == Write
}

type Access struct {
	Orders Permissions
}
