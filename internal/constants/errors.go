package constants

import "errors"

var (
	ErrorNoRowsAffected = errors.New("no rows affected")
	ErrorDataIsEmpty    = errors.New("data is empty")
)
