// Package apperrors is a custom error type that wraps an original error with a custmo message.
// It is used by the application to provide more context about errors that occur.
// Typically printed to logs and returned to clients by @code{api.WriteError}.
package apperrors

type AppError struct {
	OriginalError error
	ErrorMessage  string
}

func Wrap(originalError error, errorMessage string) *AppError {
	return &AppError{
		OriginalError: originalError,
		ErrorMessage:  errorMessage,
	}
}

func NewInternalError(originalError error, errorMessage string) *AppError {
	return Wrap(originalError, "Internal server error: "+errorMessage)
}

func (e *AppError) Error() string {
	return getErrorMessage(e)
}

func getErrorMessage(err *AppError) string {
	if err.OriginalError == nil {
		return err.ErrorMessage
	} else if err.ErrorMessage == "" {
		return err.OriginalError.Error()
	} else {
		return err.ErrorMessage + ":\n" + err.OriginalError.Error()
	}
}
