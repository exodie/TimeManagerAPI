package customErrors

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewAppError(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

func (e *AppError) Error() string {
	return e.Message
}
