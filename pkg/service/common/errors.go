package common

import (
	"errors"

	"github.com/bufbuild/connect-go"
)

var (
	errMissingRequestBody = errors.New("missing request body")
)

func CheckMessage[T any](req *connect.Request[T]) *connect.Error {
	if req.Msg == nil {
		return connect.NewError(connect.CodeInvalidArgument, errMissingRequestBody)
	}
	return nil
}
