package client

import (
	"google.golang.org/protobuf/proto"

	api "github.com/benliusf/trader_workstation_go_sdk/api/v104401"
)

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

// Access Control List
type ACL struct {
	Orders Permissions
}

type Role ACL

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

func ValidateRequestACL(r *Role, m proto.Message) error {
	switch m.(type) {
	case *api.PlaceOrderRequest:
		if !CanCreate(r.Orders) {
			return ErrNoCreate
		}
	case *api.CancelOrderRequest, *api.GlobalCancelRequest:
		if !CanDelete(r.Orders) {
			return ErrNoDelete
		}
	}
	return nil
}
