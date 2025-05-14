package quote

import "errors"

var (
	ErrNotFound       = errors.New("not Found")
	ErrInternalServer = errors.New("internel server error")
)
