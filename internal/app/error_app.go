package app

type AppError struct {
	Code    string
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

const (
	ErrBadRequest   = "BAD_REQUEST"
	ErrInvalidInput = "INVALID_INPUT"
	ErrNotFound     = "NOT_FOUND"
	ErrConflict     = "CONFLICT"
	ErrForbidden    = "FORBIDDEN"
	ErrInternal     = "INTERNAL"
)

func BadRequest(msg string) error {
	return &AppError{Code: ErrBadRequest, Message: msg}
}

func NotFound(msg string) error {
	return &AppError{Code: ErrNotFound, Message: msg}
}

func InvalidInput(msg string) error {
	return &AppError{Code: ErrInvalidInput, Message: msg}
}

func Conflict(msg string) error {
	return &AppError{Code: ErrConflict, Message: msg}
}

func Forbidden(msg string) error {
	return &AppError{Code: ErrForbidden, Message: msg}
}

func Internal(msg string) error {
	return &AppError{Code: ErrInternal, Message: msg}
}
