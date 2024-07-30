package service

import (
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var errorMapping = map[string]struct {
	code codes.Code
	msg  string
}{
	"no rows in result set":                          {codes.NotFound, "data is empty"},
	"duplicate key value violates unique constraint": {codes.AlreadyExists, "variable value is already exists"},
	"violates foreign key constraint":                {codes.InvalidArgument, "foreign key violation"},
	"no rows affected":                               {codes.NotFound, "variable value is not exists"},
}

func serviceError(err error, code codes.Code) error {
	if err == nil {
		return nil
	}

	errMsg := err.Error()

	for substr, mapping := range errorMapping {
		if strings.Contains(errMsg, substr) {
			return status.Error(mapping.code, mapping.msg)
		}
	}

	if code != codes.OK {
		return status.Error(code, errMsg)
	}

	return status.Error(codes.Unknown, errMsg)
}
