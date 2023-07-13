package apperror

import "fmt"

type AppError struct {
	ErrorMassage string
	ErrorCode    int
}

func (ae AppError) Error() string {
	return fmt.Sprintf("%v - %v", ae.ErrorCode, ae.ErrorMassage)
}

func NewAppError(errorCode int, errorMassage string) error {
	return AppError{
		ErrorCode:    errorCode,
		ErrorMassage: errorMassage,
	}
}
