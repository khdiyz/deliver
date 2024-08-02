package constants

import "errors"

var (
	ErrNoRowsAffected   = errors.New("no rows affected")
	ErrDataIsEmpty      = errors.New("data is empty")
	ErrInvalidUserId    = errors.New("invalid user id")
	ErrPermissionDenied = errors.New("permission denied")
	ErrRoleNotFound     = errors.New("role not found")
)
