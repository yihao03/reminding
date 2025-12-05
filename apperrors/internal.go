package apperrors

type InternalServerError struct{}

func (e *InternalServerError) Error() string {
	return "An internal server error occurred"
}
