package exception


type InternalServerError struct {
	Error interface{}
}

func PanicIfInternalServerError(err interface{}) {
	if err != nil {
		panic(InternalServerError{
			Error: err.(error).Error(),
		})
	}
}
