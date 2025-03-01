package exceptions

type ValidationError struct {
	Message string
}

func (validationError ValidationError) Error() string {
	return validationError.Message
}

func ValidationForm(err interface{}) {
	if err != nil {
		panic(ValidationError{
			Message: err.(error).Error(),
		})
	}
}
