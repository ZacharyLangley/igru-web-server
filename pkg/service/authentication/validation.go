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

var errPasswordContainsNonASCII = errors.New("password contains non-ascii characters")

func validatePassword(password string) error {
	if len(password) == 0 {
		return ValidationRequiredError{FieldName: "password"}
	}
	if len(password) < minPasswordLength {
		return ValidationLengthError{
			FieldName: "password",
			Minimum:   minPasswordLength,
			Current:   password,
		}
	}
	for i := 0; i < len(password); i++ {
		if password[i] >= utf8.RuneSelf {
			return errPasswordContainsNonASCII
		}
	}
	return nil
}

func validateCreateUserRequest(req *v1.CreateUserRequest) error {
	if req == nil {
		return connect.NewError(connect.CodeInternal, errInvalidRequestBody)
	}
	if _, err := mail.ParseAddress(req.Email); err != nil {
		return connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid email address: %w", err))
	}
	if err := validatePassword(req.Password); err != nil {
		return connect.NewError(connect.CodeInvalidArgument, err)
	}
	if req.FullName != "" {
		if len(req.FullName) > 255 {
			return connect.NewError(connect.CodeInvalidArgument, ValidationLengthError{
				FieldName: "request.fullName",
				Maximum:   255,
				Current:   req.FullName,
			})
		}
	}
	return nil
}

var errInvalidRequestBody = ValidationRequiredError{FieldName: "request body"}

func validateResetUserPasswordRequest(req *v1.ResetUserPasswordRequest) error {
	if req == nil {
		return connect.NewError(connect.CodeInternal, errInvalidRequestBody)
	}
	if err := validatePassword(req.Password); err != nil {
		return connect.NewError(connect.CodeInvalidArgument, err)
	}
	return nil
}

func validateDeleteUserRequest(req *v1.DeleteUserRequest) error {
	if req == nil {
		return connect.NewError(connect.CodeInternal, errInvalidRequestBody)
	}
	_, err := uuid.Parse(req.UserId)
	if err != nil {
		return connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid user id: %w", err))
	}
	return nil
}

func validateUpdateUserRequest(req *v1.UpdateUserRequest) error {
	if req == nil {
		return connect.NewError(connect.CodeInternal, errInvalidRequestBody)
	}
	if req.User == nil {
		return connect.NewError(connect.CodeInvalidArgument, ValidationRequiredError{FieldName: "request.user"})
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
			return connect.NewError(connect.CodeInvalidArgument, ValidationLengthError{
				FieldName: "request.user.fullName",
				Maximum:   255,
				Current:   req.User.FullName,
			})
		}
	}
	return nil
}

func validateCreateGroupRequest(req *v1.CreateGroupRequest) error {
	if req == nil {
		return connect.NewError(connect.CodeInternal, errInvalidRequestBody)
	}
	if len(req.Name) < 1 {
		return connect.NewError(connect.CodeInvalidArgument, ValidationRequiredError{FieldName: "request.name"})
	}
	if len(req.Name) > 255 {
		return connect.NewError(connect.CodeInvalidArgument, ValidationLengthError{
			FieldName: "request.name",
			Maximum:   255,
			Current:   req.Name,
		})
	}
	return nil
}

func validateDeleteGroupRequest(req *v1.DeleteGroupRequest) error {
	if req == nil {
		return connect.NewError(connect.CodeInternal, errInvalidRequestBody)
	}
	_, err := uuid.Parse(req.Id)
	if err != nil {
		return connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid group id: %w", err))
	}
	return nil
}

type ValidationRequiredError struct {
	FieldName string
}

func (v ValidationRequiredError) Error() string {
	return fmt.Sprintf("missing %s", v.FieldName)
}

type ValidationLengthError struct {
	FieldName string
	Minimum   int
	Maximum   int
	Current   string
}

func (v ValidationLengthError) Error() string {
	length := len(v.Current)
	if length < v.Minimum {
		return fmt.Sprintf("%s cannot be longer than %d characters long", v.FieldName, v.Minimum)
	}
	if v.Maximum != 0 && length > v.Maximum {
		return fmt.Sprintf("%s cannot be less than %d characters long", v.FieldName, v.Maximum)
	}
	return fmt.Sprintf("%s must be between %d and %d characters long", v.FieldName, v.Minimum, v.Maximum)
}

func validateUpdateGroupRequest(req *v1.UpdateGroupRequest) error {
	if req == nil {
		return connect.NewError(connect.CodeInternal, ValidationRequiredError{"request"})
	}
	if req.Group == nil {
		return connect.NewError(connect.CodeInvalidArgument, ValidationRequiredError{"request.group"})
	}
	if len(req.Group.Name) > 0 {
		return connect.NewError(connect.CodeInvalidArgument, ValidationRequiredError{"request.group.name"})
	}
	if len(req.Group.Name) > 255 {
		return connect.NewError(connect.CodeInvalidArgument, ValidationLengthError{
			FieldName: "request.group.name",
			Maximum:   255,
			Current:   req.Group.Name,
		})
	}
	return nil
}

const maxPageLength = 1000
const defaultPageLength = 100

type paginationValueError struct {
	ValueName    string
	MinimumValue int32
	MaximumValue int32
	CurrentValue int32
}

func (p paginationValueError) Error() string {
	if p.CurrentValue < p.MinimumValue {
		return fmt.Sprintf("pagination %s cannot be less than %d. got %d", p.ValueName, p.MinimumValue, p.CurrentValue)
	} else if p.MaximumValue != 0 && p.CurrentValue > p.MaximumValue {
		return fmt.Sprintf("pagination %s cannot be more than %d. got %d", p.ValueName, p.MaximumValue, p.CurrentValue)
	}
	return fmt.Sprintf("pagination %s must be between %d and %d. got %d", p.ValueName, p.MinimumValue, p.MaximumValue, p.CurrentValue)
}

type paginationRequest interface {
	GetPagination() *commonv1.PaginationRequest
}

func validatePaginationRequest(req paginationRequest) error {
	if req == nil {
		return connect.NewError(connect.CodeInternal, errInvalidRequestBody)
	}
	page := req.GetPagination()
	if page != nil {
		if page.Cursor < 0 {
			return connect.NewError(connect.CodeInvalidArgument, paginationValueError{
				ValueName:    "cursor",
				MinimumValue: 0,
				CurrentValue: page.Cursor,
			})
		}
		if page.Length < 0 || page.Length > maxPageLength {
			return connect.NewError(connect.CodeInvalidArgument, paginationValueError{
				ValueName:    "length",
				MinimumValue: 0,
				MaximumValue: maxPageLength,
				CurrentValue: page.Length,
			})
		}
		if page.Length == 0 {
			page.Length = defaultPageLength
		}
	}
	return nil
}

func validateCreateSessionRequest(req *v1.CreateSessionRequest) error {
	if req == nil {
		return connect.NewError(connect.CodeInternal, errInvalidRequestBody)
	}
	if _, err := mail.ParseAddress(req.Email); err != nil {
		return connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid email address: %w", err))
	}
	if err := validatePassword(req.Password); err != nil {
		return connect.NewError(connect.CodeInvalidArgument, err)
	}
	return nil
}
