package response

type Status struct {
	Code        int    `json:"code"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

var (
	OK = Status{
		Code:        200,
		Status:      "OK",
		Description: "The request has succeeded.",
	}
	Created = Status{
		Code:        201,
		Status:      "CREATED",
		Description: "The request has been fulfilled and has resulted in one or more new resources being created.",
	}
	NoContent = Status{
		Code:        204,
		Status:      "NO_CONTENT",
		Description: "There is no content to send for this request, but the headers may be useful.",
	}
	BadEnvironment = Status{
		Code:        400,
		Status:      "BAD_ENVIRONMENT",
		Description: "The service has an invalid environment value.",
	}
	BadRequest = Status{
		Code:        400,
		Status:      "BAD_REQUEST",
		Description: "The server could not understand the request due to invalid syntax.",
	}
	InvalidArgument = Status{
		Code:        400,
		Status:      "INVALID_ARGUMENT",
		Description: "Incorrect request parameters specified. Details are provided in the details field.",
	}
	FailedPrecondition = Status{
		Code:        400,
		Status:      "FAILED_PRECONDITION",
		Description: "The operation was canceled because the conditions required for the operation were not met.",
	}
	OutOfRange = Status{
		Code:        400,
		Status:      "OUT_OF_RANGE",
		Description: "Out of range. For example, searching or reading outside of the file.",
	}
	Unauthorized = Status{
		Code:        401,
		Status:      "UNAUTHORIZED",
		Description: "The operation requires authentication.",
	}
	Forbidden = Status{
		Code:        403,
		Status:      "FORBIDDEN",
		Description: "...",
	}
	PermissionDenied = Status{
		Code:        403,
		Status:      "PERMISSION_DENIED",
		Description: "The user has no permissions required to perform the operation.",
	}
	NotFound = Status{
		Code:        404,
		Status:      "NOT_FOUND",
		Description: "The requested resource not found.",
	}
	AlreadyExists = Status{
		Code:        409,
		Status:      "ALREADY_EXISTS",
		Description: "Arguments already exists.",
	}
	Aborted = Status{
		Code:        409,
		Status:      "ABORTED",
		Description: "The operation failed due to a concurrent computing conflict, such as an invalid sequence of commands or an aborted transaction.",
	}
	FailedDependency = Status{
		Code:        424,
		Status:      "FAILED_DEPENDENCY",
		Description: "The operation failed due to a concurrent computing conflict, such as an invalid sequence of commands or an aborted transaction.",
	}
	TooManyRequests = Status{
		Code:        429,
		Status:      "TOO_MANY_REQUESTS",
		Description: "The user has sent too many requests in a given amount of time.",
	}
	ResourceExhausted = Status{
		Code:        429,
		Status:      "RESOURCE_EXHAUSTED",
		Description: "The request limit exceeded.",
	}
	Canceled = Status{
		Code:        499,
		Status:      "CANCELED",
		Description: "The operation was aborted on the client side.",
	}
	Internal = Status{
		Code:        500,
		Status:      "INTERNAL",
		Description: "Internal server error. This error means that the operation cannot be performed due to a server-side technical problem. For example, due to insufficient computing resources.",
	}
	GRPCError = Status{
		Code:        500,
		Status:      "GRPC_ERROR",
		Description: "The gRPC request failed.",
	}
	Unknown = Status{
		Code:        500,
		Status:      "UNKNOWN",
		Description: "Unknown error.",
	}
	DataLoss = Status{
		Code:        500,
		Status:      "DATA_LOSS",
		Description: "Permanent data loss or damage.",
	}
	Unimplemented = Status{
		Code:        501,
		Status:      "UNIMPLEMENTED",
		Description: "The operation is not supported by the service.",
	}
	Unavailable = Status{
		Code:        503,
		Status:      "UNAVAILABLE",
		Description: "The service is currently unavailable. Try again in a few seconds.",
	}
	DeadlineExceeded = Status{
		Code:        504,
		Status:      "DEADLINE_EXCEEDED",
		Description: "Exceeded the server response timeout.",
	}
)