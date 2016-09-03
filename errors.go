package rememberme

import (
	"errors"
)

var (
	ErrorBadRequest     = errors.New("bad request")
	ErrorSessionExpired = errors.New("session expired")
)
