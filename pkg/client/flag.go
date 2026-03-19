package client

type Permissions uint64

const (
	Read Permissions = 1 << iota
	Create
	Delete
)

func CanRead(o Permissions) bool {
	return (o & Read) == Read
}

func CanCreate(o Permissions) bool {
	return (o & Create) == Create
}

func CanDelete(o Permissions) bool {
	return (o & Delete) == Delete
}

type Access struct {
	Orders Permissions
}

type Role Access

func ReadOnly() Role {
	return Role{
		Orders: Read,
	}
}

func ReadAndWrite() Role {
	return Role{
		Orders: Read | Create | Delete,
	}
}
