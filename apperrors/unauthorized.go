package apperrors

type UnauthorizedError struct{}

func (e *UnauthorizedError) Error() string {
	return "User is unauthorized to perform this action"
}
