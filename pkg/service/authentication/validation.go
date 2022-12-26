package authentication

import (
	"errors"
	"fmt"
	"net/mail"
	"unicode/utf8"

	v1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1"
	commonv1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/common/v1"
	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
)

const minPasswordLength = 8

func validatePassword(password string) error {
	if len(password) == 0 {
		return errors.New("missing password")
	}
	if len(password) < minPasswordLength {
		return fmt.Errorf("password is too short. needs to be at least %d characters long", minPasswordLength)
	}
	for i := 0; i < len(password); i++ {
		if password[i] >= utf8.RuneSelf {
			return errors.New("password contains non-ascii characters")
		}
	}
	return nil
}

func validateCreateUserRequest(req *v1.CreateUserRequest) error {
	if req == nil {
		return connect.NewError(connect.CodeInternal, errors.New("missing request body"))
	}
	if _, err := mail.ParseAddress(req.Email); err != nil {
		return connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid email address: %w", err))
	}
	if err := validatePassword(req.Password); err != nil {
		return connect.NewError(connect.CodeInvalidArgument, err)
	}
	if req.FullName != "" {
		if len(req.FullName) > 255 {
			return connect.NewError(connect.CodeInvalidArgument, errors.New("full name cannot be longer than 255 characters long."))
		}
	}
	return nil
}

func validateResetUserPasswordRequest(req *v1.ResetUserPasswordRequest) error {
	if req == nil {
		return connect.NewError(connect.CodeInternal, errors.New("missing request body"))
	}
	if err := validatePassword(req.Password); err != nil {
		return connect.NewError(connect.CodeInvalidArgument, err)
	}
	return nil
}

func validateDeleteUserRequest(req *v1.DeleteUserRequest) error {
	if req == nil {
		return connect.NewError(connect.CodeInternal, errors.New("missing request body"))
	}
	_, err := uuid.Parse(req.UserId)
	if err != nil {
		return connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid user id: %w", err))
	}
	return nil
}

func validateUpdateUserRequest(req *v1.UpdateUserRequest) error {
	if req == nil {
		return connect.NewError(connect.CodeInternal, errors.New("missing request body"))
	}
	if req.User == nil {
		return connect.NewError(connect.CodeInvalidArgument, errors.New("missing user object"))
	}
	_, err := uuid.Parse(req.User.Id)
	if err != nil {
		return connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid user id: %w", err))
	}
	if _, err := mail.ParseAddress(req.User.Email); err != nil {
		return connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid email address: %w", err))
	}
	if req.User.FullName != "" {
		if len(req.User.FullName) > 255 {
			return connect.NewError(connect.CodeInvalidArgument, errors.New("full name cannot be longer than 255 characters long."))
		}
	}
	return nil
}

func validateCreateGroupRequest(req *v1.CreateGroupRequest) error {
	if req == nil {
		return connect.NewError(connect.CodeInternal, errors.New("missing request body"))
	}
	if len(req.Name) > 0 {
		return connect.NewError(connect.CodeInvalidArgument, errors.New("name is required"))
	}
	if len(req.Name) > 255 {
		return connect.NewError(connect.CodeInvalidArgument, errors.New("name cannot be longer than 255 characters long."))
	}
	return nil
}

func validateDeleteGroupRequest(req *v1.DeleteGroupRequest) error {
	if req == nil {
		return connect.NewError(connect.CodeInternal, errors.New("missing request body"))
	}
	_, err := uuid.Parse(req.Id)
	if err != nil {
		return connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid group id: %w", err))
	}
	return nil
}

func validateUpdateGroupRequest(req *v1.UpdateGroupRequest) error {
	if req == nil {
		return connect.NewError(connect.CodeInternal, errors.New("missing request body"))
	}
	if req.Group == nil {
		return connect.NewError(connect.CodeInvalidArgument, errors.New("missing group object"))
	}
	if len(req.Group.Name) > 0 {
		return connect.NewError(connect.CodeInvalidArgument, errors.New("name is required"))
	}
	if len(req.Group.Name) > 255 {
		return connect.NewError(connect.CodeInvalidArgument, errors.New("name cannot be longer than 255 characters long."))
	}
	return nil
}

const maxPageLength = 1000
const defaultPageLength = 100

type paginationRequest interface {
	GetPagination() *commonv1.PaginationRequest
}

func validatePaginationRequest(req paginationRequest) error {
	if req == nil {
		return connect.NewError(connect.CodeInternal, errors.New("missing request body"))
	}
	page := req.GetPagination()
	if page != nil {
		if page.Cursor < 0 {
			return connect.NewError(connect.CodeInvalidArgument, errors.New("pagination cursor cannot be less than 0"))
		}
		if page.Length < 0 {
			return connect.NewError(connect.CodeInvalidArgument, errors.New("pagination length cannot be less than 0"))
		}
		if page.Length > maxPageLength {
			return connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("pagination length cannot be more than %d", maxPageLength))
		}
		if page.Length == 0 {
			page.Length = defaultPageLength
		}
	}
	return nil
}

func validateGetTokenRequest(req *v1.GetTokenRequest) error {
	if req == nil {
		return connect.NewError(connect.CodeInternal, errors.New("missing request body"))
	}
	if _, err := mail.ParseAddress(req.Email); err != nil {
		return connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid email address: %w", err))
	}
	if err := validatePassword(req.Password); err != nil {
		return connect.NewError(connect.CodeInvalidArgument, err)
	}
	return nil
}
