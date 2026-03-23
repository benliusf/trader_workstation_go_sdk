package client

import "errors"

var ErrClientNotReady = errors.New("client not ready")

var ErrNotAllowed = errors.New("not allowed")

var ErrInvalidParam = errors.New("invalid param")
