package exception

type Duplicate struct {
	Error interface{}
}

func PanicDuplicate(idSeller int, msg string) {
	if idSeller > 0 {
		panic(Duplicate{
			Error: msg,
		})
	}
}