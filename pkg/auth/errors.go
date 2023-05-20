package auth

import "errors"

var ErrMissingToken = errors.New("missing auth token")
var ErrPermissionDenied = errors.New("permission denied")
