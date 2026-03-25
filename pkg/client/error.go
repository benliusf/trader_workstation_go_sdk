package client

import (
	"errors"
	"fmt"
)

var ErrClientNotReady = errors.New("client not ready")

var ErrNotAllowed = errors.New("not allowed")

var ErrNoCreate = fmt.Errorf("%w: no create permission", ErrNotAllowed)

var ErrNoDelete = fmt.Errorf("%w: no delete permission", ErrNotAllowed)

var ErrInvalidParam = errors.New("invalid param")
