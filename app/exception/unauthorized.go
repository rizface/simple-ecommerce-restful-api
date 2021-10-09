package exception

type Unauthorized struct {
	Error interface{}
}

func PanicUnauthorized(err interface{}) {
	if err != nil {
		panic(Unauthorized{
			Error:err.(error).Error(),
		})
	}
}
